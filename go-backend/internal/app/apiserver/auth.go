package apiserver

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-next-pizza/internal/app/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (s *server) handleRegister() http.HandlerFunc {
	type request struct {
		Email    string
		Password string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}

		if _, err := s.storage.User().CreateUser(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) handleSessionCreate() http.HandlerFunc {
	type request struct {
		Email    string
		Password string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.storage.User().FindByEmail(req.Email)

		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, ErrIncorrectEmailOrPassword)
			return
		}

		access := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": u.Id,
			"exp": time.Now().Add(time.Duration(s.jwtTTLSeconds) * time.Second).Unix(),
			"iat": time.Now().Unix(),
			"jti": uuid.New().String(),
		})

		refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": u.Id,
			"exp": time.Now().Add(time.Duration(s.refreshTTLSeconds) * time.Second).Unix(),
			"iat": time.Now().Unix(),
			"typ": "refresh",
			"jti": uuid.New().String(),
		})

		accessToken, err := access.SignedString(s.jwtSecret)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		refreshToken, err := refresh.SignedString(s.jwtSecret)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, map[string]string{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		})
	}
}

func (s *server) handleSessionRefresh() http.HandlerFunc {
	type request struct {
		RefreshToken string `json:"refreshToken"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req request

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.RefreshToken == "" {
			s.error(w, r, http.StatusBadRequest, ErrIncorrectEmailOrPassword)
			return
		}

		token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}

			return s.jwtSecret, nil
		})

		if err != nil || !token.Valid {
			s.error(w, r, http.StatusUnauthorized, ErrNotAuthenticated)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || claims["typ"] != "refresh" {
			s.error(w, r, http.StatusUnauthorized, ErrNotAuthenticated)
			return
		}

		userIDFloat, ok := claims["sub"].(float64)

		if !ok {
			s.error(w, r, http.StatusUnauthorized, ErrNotAuthenticated)
			return
		}

		if _, err := s.storage.User().FindById(int(userIDFloat)); err != nil {
			s.error(w, r, http.StatusUnauthorized, ErrNotAuthenticated)
			return
		}

		access := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": int(userIDFloat),
			"exp": time.Now().Add(time.Duration(s.jwtTTLSeconds) * time.Second).Unix(),
			"iat": time.Now().Unix(),
		})

		accessToken, err := access.SignedString(s.jwtSecret)

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, map[string]string{
			"accessToken": accessToken,
		})
	}
}


func (s *server) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			s.error(w, r, http.StatusUnauthorized, ErrNotAuthenticated)
			return
		}

		const bearer = "Bearer "
		if len(auth) <= len(bearer) || auth[:len(bearer)] != bearer {
			s.error(w, r, http.StatusUnauthorized, ErrNotAuthenticated)
			return
		}

		tokenStr := auth[len(bearer):]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return s.jwtSecret, nil
		})

		if err != nil || !token.Valid {
			s.error(w, r, http.StatusUnauthorized, ErrNotAuthenticated)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			s.error(w, r, http.StatusUnauthorized, ErrNotAuthenticated)
			return
		}

		userIDFloat, ok := claims["sub"].(float64)
		if !ok {
			s.error(w, r, http.StatusUnauthorized, ErrNotAuthenticated)
			return
		}

		u, err := s.storage.User().FindById(int(userIDFloat))
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, ErrNotAuthenticated)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
}
