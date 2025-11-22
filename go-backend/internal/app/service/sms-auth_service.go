package service

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"sync"
	"time"

	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/storage"
	"golang.org/x/time/rate"
)

const (
	CODE_LENGTH = 6
	NOTISEND_API_KEY = "8789aa6a751ec66aeffccfc1bed5cc21"
	NOTISEND_PROJECT = "go-next-pizza"
	NOTISEND_API_URL = "https://sms.notisend.ru/api/message/send"
)

type SMSAuthService struct {
	userRepo storage.UserRepository
  smsCodeRepo storage.SMSCodeRepository
}

type SMSAuthToken struct {
	Code string `json:"code"`
	ExpiresIn int `json:"expires_in"`
}

func NewSMSAuthService(userRepo storage.UserRepository) *SMSAuthService {
	return &SMSAuthService{
      userRepo: userRepo,
	}
}

func NewSMSAuthServiceWithCodes(userRepo storage.UserRepository, smsCodeRepo storage.SMSCodeRepository) *SMSAuthService {
    return &SMSAuthService{
      userRepo:    userRepo,
      smsCodeRepo: smsCodeRepo,
    }
}

func (sas *SMSAuthService) Register(user *model.User) (*model.User, error) {
	if err := user.Validate(); err != nil {
		log.Printf("Пользователь не прошел валидацию");
		return nil, err
	}

	createdUser, err := sas.userRepo.CreateUser(user)

	if err != nil {
		log.Printf("Не удалось создать пользователя")
		return nil, err
	}

	return createdUser, nil
}

func (sas *SMSAuthService) Login(phone string) (*SMSAuthToken, error) {
	user, err := sas.userRepo.FindByPhone(phone)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return &SMSAuthToken{
		Code: sas.generateCode(CODE_LENGTH),
		ExpiresIn: 60,
	}, nil
}

func (sas *SMSAuthService) SendCode(phone string, ttl time.Duration) (string, error) {
    if sas.smsCodeRepo == nil {
        return "", errors.New("sms code repository is not configured")
    }
    code := sas.generateCode(CODE_LENGTH)
    sum := sha256.Sum256([]byte(code))
    hash := hex.EncodeToString(sum[:])
    expiresAt := time.Now().Add(ttl).Unix()
    if err := sas.smsCodeRepo.SaveCode(phone, hash, expiresAt); err != nil {
        return "", err
    }
    return code, nil
}

func (sas *SMSAuthService) VerifyCode(phone string, code string) (bool, error) {
    if sas.smsCodeRepo == nil {
        return false, errors.New("sms code repository is not configured")
    }

    storedHash, expiresAt, err := sas.smsCodeRepo.GetLatestCodeHash(phone)
    if err != nil {
        return false, err
    }

    if time.Now().Unix() > expiresAt {
        _ = sas.smsCodeRepo.DeleteCodes(phone)
        return false, errors.New("code expired")
    }

    sum := sha256.Sum256([]byte(code))
    inputHash := hex.EncodeToString(sum[:])
    if inputHash != storedHash {
        return false, errors.New("invalid code")
    }

    _ = sas.smsCodeRepo.DeleteCodes(phone)
    return true, nil
}

func (sas *SMSAuthService) generateCode(length int) string {
	var numbers = []rune("0123456789")

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, length)
	for i := range b {
		b[i] = numbers[rng.Intn(len(numbers))]
	}
	return string(b)
}

// GetUserByPhone provides access for handlers to fetch user after verification
func (sas *SMSAuthService) GetUserByPhone(phone string) (*model.User, error) {
    return sas.userRepo.FindByPhone(phone)
}

func (sas *SMSAuthService) GetUserByEmail(email string) (*model.User, error) {
    return sas.userRepo.FindByEmail(email)
}

var (
	limiter   = rate.NewLimiter(1/60.0, 1)
	mu        sync.Mutex
	codeStore = map[string]struct {
		code      string
		expiresAt time.Time
	}{}
)

type smsSendRequest struct {
	Phone string `json:"phone"`
}

type smsVerifyRequest struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

func handleSMS(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	mu.Lock()
	if !limiter.Allow() {
		mu.Unlock()
		http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
		return
	}
	
	mu.Unlock()

	var req smsSendRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Phone == "" {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	code := (&SMSAuthService{}).generateCode(CODE_LENGTH)
	mu.Lock()
	codeStore[req.Phone] = struct {
		code      string
		expiresAt time.Time
	}{code: code, expiresAt: time.Now().Add(2 * time.Minute)}
	mu.Unlock()

	message := fmt.Sprintf("Your verification code is %s", code)
	if err := sendSMS(req.Phone, message); err != nil {
		log.Printf("Ошибка отправки SMS: %v", err)
		http.Error(w, fmt.Sprintf("Failed to send SMS: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"sent": true, "code": code, "expires_in": 120})
}

func handleVerify(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req smsVerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Phone == "" || req.Code == "" {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	mu.Lock()
	entry, ok := codeStore[req.Phone]
	mu.Unlock()
	if !ok {
		http.Error(w, "code not found", http.StatusBadRequest)
		return
	}
	if time.Now().After(entry.expiresAt) {
		mu.Lock()
		delete(codeStore, req.Phone)
		mu.Unlock()
		http.Error(w, "code expired", http.StatusBadRequest)
		return
	}
	if req.Code != entry.code {
		http.Error(w, "invalid code", http.StatusUnauthorized)
		return
	}

	mu.Lock()
	delete(codeStore, req.Phone)
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"verified": true})
}

func sendSMS(destination, message string) error {
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	writer.WriteField("project", NOTISEND_PROJECT)
	writer.WriteField("recipients", destination)
	writer.WriteField("message", message)
	writer.WriteField("apikey", NOTISEND_API_KEY)

	if err := writer.Close(); err != nil {
		return fmt.Errorf("failed to close multipart writer: %w", err)
	}

	req, err := http.NewRequest("POST", NOTISEND_API_URL, &requestBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	contentType := writer.FormDataContentType()
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Accept", "application/json")
	
	log.Printf("Отправка SMS: project=%s, recipients=%s, message=%s", NOTISEND_PROJECT, destination, message)
	log.Printf("Content-Type: %s", contentType)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("network error: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("API вернул ошибку: статус %d, тело ответа: %s", resp.StatusCode, string(body))
		return fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
	}

	var responseData map[string]any
	if err := json.Unmarshal(body, &responseData); err == nil {
		if status, ok := responseData["status"].(string); ok && status != "success" {
			log.Printf("API вернул неуспешный статус: %s, ответ: %s", status, string(body))
			return fmt.Errorf("provider returned non-success status: %s", string(body))
		}
	}

	log.Printf("Ответ провайдера: %s", string(body))
	log.Printf("SMS успешно отправлено на номер %s", destination)
	return nil
}

func rateLimit(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		allowed := limiter.Allow()
		mu.Unlock()
		if !allowed {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	}
}