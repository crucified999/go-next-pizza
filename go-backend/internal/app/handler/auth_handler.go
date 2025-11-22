package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/service"
	"github.com/go-next-pizza/internal/app/storage"
)

type AuthHandler struct {
	authService *service.AuthService
	userRepo    storage.UserRepository
}

type RegisterRequest struct {
	Phone    string `json:"phone"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CheckAuthResponse struct {
	// User *model.User `json:"user"`
	Id  int `json:"id"`
	Phone string `json:"phone"`
	Name string `json:"name"`
	Email string `json:"email"`
	Authenticated bool `json:"authenticated"`
}


func NewAuthHandler(authService *service.AuthService, userRepo storage.UserRepository) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		userRepo:    userRepo,
	}
}

func (ah *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}


	user := &model.User{
		Phone:    req.Phone,
	}

	createdUser, err := ah.authService.Register(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdUser)
}

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	tokens, err := ah.authService.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokens)
}

func (ah *AuthHandler) CheckAuth(w http.ResponseWriter, r *http.Request) {
	var tokenString string
	var res CheckAuthResponse

	// Проверяем access token из разных источников
	if cookie, err := r.Cookie("access"); err == nil && cookie != nil {
		tokenString = cookie.Value
		log.Printf("Found access token in cookie")
	} else if authHeader := r.Header.Get("Authorization"); authHeader != "" {
		tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			log.Printf("Invalid Authorization header format")
			tokenString = ""
		} else {
			log.Printf("Found access token in Authorization header")
		}
	}

	// Пытаемся использовать access token
	if tokenString != "" {
		userID, err := ah.authService.ValidateToken(tokenString)
		if err == nil {
			user, err := ah.userRepo.FindById(userID)
			if err == nil {
				log.Printf("Access token valid for user ID: %d", userID)
				res.Id = userID
				res.Phone = user.Phone

				if (user.Email.Valid) {
					res.Email = user.Email.String
				} else {
					res.Email = ""
				}

				if (user.Name.Valid) {
					res.Name = user.Name.String
				} else {
					res.Name = ""
				}

				res.Authenticated = true
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(res)
				return
			} else {
				log.Printf("User not found for ID: %d, error: %v", userID, err)
			}
		} else {
			log.Printf("Access token invalid: %v", err)
		}
	} else {
		log.Printf("No access token provided")
	}

	// Пытаемся обновить токены через refresh token
	refreshCookie, err := r.Cookie("refresh")
	if err == nil && refreshCookie != nil {
		log.Printf("Attempting token refresh")
		
		userID, err := ah.authService.ValidateToken(refreshCookie.Value)
		if err == nil {
			tokens, err := ah.authService.GenerateTokensForUserID(userID)
			if err == nil {
				// Устанавливаем новые cookies
				http.SetCookie(w, &http.Cookie{
					Name:     "access",
					Value:    tokens.AccessToken,
					Path:     "/",
					HttpOnly: true,
					Secure:   true,
					SameSite: http.SameSiteLaxMode,
					MaxAge:   tokens.ExpiresIn,
				})

				http.SetCookie(w, &http.Cookie{
					Name:     "refresh",
					Value:    tokens.RefreshToken,
					Path:     "/",
					HttpOnly: true,
					Secure:   true,
					SameSite: http.SameSiteLaxMode,
					MaxAge:   60 * 60 * 24 * 180,
				})

				user, err := ah.userRepo.FindById(userID)
				if err == nil {
					log.Printf("Tokens refreshed successfully for user ID: %d", userID)
					res.Id = userID
					res.Phone = user.Phone

				if (user.Email.Valid) {
					res.Email = user.Email.String
				} else {
					res.Email = ""
				}

				if (user.Name.Valid) {
					res.Name = user.Name.String
				} else {
					res.Name = ""
				}
					res.Authenticated = true
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(res)
					return
				} else {
					log.Printf("User not found after token refresh: %v", err)
				}
			} else {
				log.Printf("Failed to generate new tokens: %v", err)
			}
		} else {
			log.Printf("Refresh token invalid: %v", err)
		}
	} else {
		log.Printf("No refresh token provided: %v", err)
	}

	// Не аутентифицирован
	log.Printf("User not authenticated")
	res.Phone = ""
	res.Email = ""
	res.Name = ""
	res.Authenticated = false
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// func (ah *AuthHandler) CheckAuth(w http.ResponseWriter, r *http.Request) {
// 	var tokenString string
// 	var res CheckAuthResponse

// 	// Проверяем сначала cookie, потом заголовок Authorization
// 	if cookie, err := r.Cookie("access"); err == nil && cookie != nil {
// 		tokenString = cookie.Value
// 	} else {
// 		authHeader := r.Header.Get("Authorization")
// 		if authHeader != "" {
// 			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
// 			if tokenString == authHeader {
// 				http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
// 				return
// 			}
// 		}
// 	}

// 	if tokenString == "" {
// 		http.Error(w, "Authorization required", http.StatusUnauthorized)
// 		return
// 	}

// 	userID, err := ah.authService.ValidateToken(tokenString)
// 	if err == nil {
// 		user, err := ah.userRepo.FindById(userID)
// 			if err == nil {
// 				res.User = user
// 				res.Authenticated = true
// 				w.Header().Set("Content-Type", "application/json")
// 				json.NewEncoder(w).Encode(res)
// 				return
// 			}
// 	}

// 	refreshCookie, err := r.Cookie("refresh")
// 	if err == nil && refreshCookie != nil {
// 		// Пытаемся обновить токены используя refresh token
// 		userID, err := ah.authService.ValidateRefreshToken(refreshCookie.Value)
// 		if err == nil {
// 			// Генерируем новые токены
// 			tokens, err := ah.authService.GenerateTokensForUserID(userID)
// 			if err == nil {
// 				// Устанавливаем новые cookies
// 				http.SetCookie(w, &http.Cookie{
// 					Name:     "access",
// 					Value:    tokens.AccessToken,
// 					Path:     "/",
// 					HttpOnly: true,
// 					Secure:   true,
// 					SameSite: http.SameSiteLaxMode,
// 					MaxAge:   tokens.ExpiresIn,
// 				})

// 				http.SetCookie(w, &http.Cookie{
// 					Name:     "refresh",
// 					Value:    tokens.RefreshToken,
// 					Path:     "/",
// 					HttpOnly: true,
// 					Secure:   true,
// 					SameSite: http.SameSiteLaxMode,
// 					MaxAge:   60 * 60 * 24, // 24 часа
// 				})

// 				// Возвращаем пользователя
// 				user, err := ah.userRepo.FindById(userID)
// 				if err == nil {
// 					res.User = user
// 					res.Authenticated = true
// 					w.Header().Set("Content-Type", "application/json")
// 					json.NewEncoder(w).Encode(res)
// 					return
// 				}
// 			}
// 		}
// 	}

// 	user, err := ah.userRepo.FindById(userID)
// 	if err != nil {
// 		http.Error(w, "User not found", http.StatusNotFound)
// 		res.User = nil
// 		res.Authenticated = false
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(res)
// 		return
// 	}

// 	res.User = user
// 	res.Authenticated = true

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(res)
// }

func (ah *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshCookie, err := r.Cookie("refresh")
	if err != nil {
			http.Error(w, "Refresh token required", http.StatusUnauthorized)
			return
	}

	// Валидируем refresh token
	userID, err := ah.authService.ValidateToken(refreshCookie.Value)
	if err != nil {
			http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
			return
	}

	// Генерируем новые токены
	tokens, err := ah.authService.GenerateTokensForUserID(userID)
	if err != nil {
			http.Error(w, "Failed to generate tokens", http.StatusInternalServerError)
			return
	}

	// Устанавливаем новые cookies
	http.SetCookie(w, &http.Cookie{
			Name:     "access",
			Value:    tokens.AccessToken,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
			MaxAge:   tokens.ExpiresIn,
	})

	http.SetCookie(w, &http.Cookie{
			Name:     "refresh",
			Value:    tokens.RefreshToken,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
			MaxAge:   60 * 60 * 24 * 180,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
			"refreshed": true,
			"expires_in": tokens.ExpiresIn,
	})
}