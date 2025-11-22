package middleware

import (
	"context"
	"net/http"

	"github.com/go-next-pizza/internal/app/service"
)

type AuthMiddleware struct {
	authService *service.AuthService
}

func NewAuthMiddleware(authService *service.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

func (am *AuthMiddleware) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	
		accessToken, err := r.Cookie("access")

		if err != nil {
			http.Error(w, "Authorization required", http.StatusUnauthorized)
			return
		}

		userID, err := am.authService.ValidateToken(accessToken.Value)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}