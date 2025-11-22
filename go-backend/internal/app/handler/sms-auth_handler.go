package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/service"
)

func normalizePhone(s string) string {
    s = strings.TrimSpace(s)
    b := make([]rune, 0, len(s))
    for i, r := range s {
        if r >= '0' && r <= '9' { b = append(b, r); continue }
        if r == '+' && i == 0 { b = append(b, r); continue }
    }
    cleaned := string(b)
    if cleaned == "" { return cleaned }
    if cleaned[0] != '+' {
        cleaned = "+" + cleaned
    }
    return cleaned
}

const (
    notisendAPIKey = "8789aa6a751ec66aeffccfc1bed5cc21"
    notisendProject = "go-next-pizza"
    notisendAPIURL = "https://sms.notisend.ru/api/message/send"
)

type SMSAuthHandler struct {
    smsAuthService *service.SMSAuthService
    authService    *service.AuthService
}

type smsSendRequest struct {
    Phone string `json:"phone"`
}

type smsVerifyRequest struct {
    Phone string `json:"phone"`
    Code  string `json:"code"`
}

func NewSMSAuthHandler(smsAuthService *service.SMSAuthService, authService *service.AuthService) *SMSAuthHandler {
    return &SMSAuthHandler{
        smsAuthService: smsAuthService,
        authService:    authService,
    }
}

func (sah *SMSAuthHandler) SendCode(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req smsSendRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Phone == "" {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    phone := normalizePhone(req.Phone)

    code, err := sah.smsAuthService.SendCode(phone, 2*time.Minute)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    message := fmt.Sprintf("Your verification code is %s", code)
    if err := sendSMS(phone, message); err != nil {
        log.Printf("Ошибка отправки SMS: %v", err)
        http.Error(w, fmt.Sprintf("Failed to send SMS: %v", err), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(map[string]any{"sent": true, "expires_in": 120})
}

func (sah *SMSAuthHandler) VerifyCode(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req smsVerifyRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Phone == "" || req.Code == "" {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    phone := normalizePhone(req.Phone)
    ok, err := sah.smsAuthService.VerifyCode(phone, req.Code)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if !ok {
        http.Error(w, "invalid code", http.StatusUnauthorized)
        return
    }

    user, err := sah.smsAuthService.GetUserByPhone(phone)

		if err != nil {
			log.Printf("Не удалось найти пользователя с таким телефоном %s", phone)
		}

    if user == nil || user.Id == 0 {
				log.Printf("Создание пользователя с телефоном %s", phone)
        masked := phone

        if len(masked) >= 4 {
            masked = masked[len(masked)-4:]
        }

        newUser := &model.User{
            Phone:   phone,
        }
				
        user, err = sah.smsAuthService.Register(newUser)
        if err != nil {
            http.Error(w, "failed to create user", http.StatusInternalServerError)
            return
        }
    }

    tokens, err := sah.authService.GenerateTokensForUserID(user.Id)
    if err != nil {
        http.Error(w, "failed to issue tokens", http.StatusInternalServerError)
        return
    }

    http.SetCookie(w, &http.Cookie{Name: "access", Value: tokens.AccessToken, Path: "/", HttpOnly: true, Secure: true, SameSite: http.SameSiteLaxMode, MaxAge: tokens.ExpiresIn})
    http.SetCookie(w, &http.Cookie{Name: "refresh", Value: tokens.RefreshToken, Path: "/", HttpOnly: true, Secure: true, SameSite: http.SameSiteLaxMode, MaxAge: 60 * 60 * 24 * 180})

    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(map[string]any{"verified": true, "expires_in": tokens.ExpiresIn, "token_type": tokens.TokenType})
}

func sendSMS(destination, message string) error {
    var requestBody bytes.Buffer
    writer := multipart.NewWriter(&requestBody)

    writer.WriteField("project", notisendProject)
    writer.WriteField("recipients", destination)
    writer.WriteField("message", message)
    writer.WriteField("apikey", notisendAPIKey)

    if err := writer.Close(); err != nil {
        return fmt.Errorf("failed to close multipart writer: %w", err)
    }

    req, err := http.NewRequest("POST", notisendAPIURL, &requestBody)
    if err != nil {
        return fmt.Errorf("failed to create request: %w", err)
    }

    contentType := writer.FormDataContentType()
    req.Header.Set("Content-Type", contentType)
    req.Header.Set("Accept", "application/json")

    log.Printf("Отправка SMS: project=%s, recipients=%s, message=%s", notisendProject, destination, message)
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