package apiserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/storage"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
    ctxKeyUser ctxKey = iota
    ctxKeyRequestId ctxKey = iota
)

type ctxKey int8

type server struct {
	router *mux.Router
	logger *slog.Logger
	storage storage.Storage
  	jwtSecret []byte
  	jwtTTLSeconds int
  	refreshTTLSeconds int
}

func newServer(storage storage.Storage, jwtSecret []byte, jwtTTLSeconds int, refreshTTLSeconds int) *server {
	logger, err := setUpLogger()

	if err != nil {
		panic(err)
	}

	s := &server{
		router: mux.NewRouter(),
		logger: logger,
		storage: storage,
    jwtSecret: jwtSecret,
    jwtTTLSeconds: jwtTTLSeconds,
  	refreshTTLSeconds: refreshTTLSeconds,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func setUpLogger() (*slog.Logger, error){
	file, err := os.OpenFile("apiserver_log.txt", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0666)

	if err != nil {
		return nil, err
	}

	logger := slog.New(slog.NewTextHandler(file, &slog.HandlerOptions{
		Level: slog.LevelDebug,	
	}))

	return logger, nil
}

func (s *server) configureRouter() {
	s.router.Use(s.setRequestId)
	s.router.Use(s.logRequest)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
    s.router.HandleFunc("/sessions", s.handleSessionCreate()).Methods("POST")
    s.router.HandleFunc("/sessions/refresh", s.handleSessionRefresh()).Methods("POST")

    private := s.router.PathPrefix("/private").Subrouter()
    private.Use(s.authenticateUser)
	private.HandleFunc("/whoami", s.handleWhoAmI()).Methods("GET")
}

func (s *server) handleWhoAmI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*model.User))
	}
}

func (s *server) setRequestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()

		w.Header().Set("X-Request-ID", id)

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestId, id)))
	}) 
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info(
			fmt.Sprintf("started %s %s", r.Method, r.RequestURI),
				slog.String("remote_addr", r.RemoteAddr),
				slog.Any("request_id", r.Context().Value(ctxKeyRequestId),

		))

        start := time.Now()
        rw := &responseWriter{ResponseWriter: w, code: http.StatusOK}

        next.ServeHTTP(rw, r)

		s.logger.Info(fmt.Sprintf("completed with %d %s in %v", rw.code, http.StatusText(rw.code), time.Since(start)))
	})
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Email string
		Password string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Email: req.Email,
			Password: req.Password,
		}

		if err := s.storage.User().CreateUser(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) handleSessionCreate() http.HandlerFunc {
	type request struct {
		Email string
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
            "accessToken": accessToken,
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

func (s *server) error(w http.ResponseWriter, r *http.Request, status int, err error) {
    s.respond(w, r, status, map[string]string{"message": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, _ *http.Request, status int, data interface{}) {
	w.WriteHeader(status)

	if data != nil {
		json.NewEncoder(w).Encode(data)
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