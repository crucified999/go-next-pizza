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
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	ctxKeyUser      ctxKey = iota
	ctxKeyRequestId ctxKey = iota
)

type ctxKey int8

type server struct {
	router            *mux.Router
	logger            *slog.Logger
	storage           storage.Storage
	jwtSecret         []byte
	jwtTTLSeconds     int
	refreshTTLSeconds int
}

func newServer(storage storage.Storage, jwtSecret []byte, jwtTTLSeconds int, refreshTTLSeconds int) *server {
	logger, err := setUpLogger()

	if err != nil {
		panic(err)
	}

	s := &server{
		router:            mux.NewRouter(),
		logger:            logger,
		storage:           storage,
		jwtSecret:         jwtSecret,
		jwtTTLSeconds:     jwtTTLSeconds,
		refreshTTLSeconds: refreshTTLSeconds,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func setUpLogger() (*slog.Logger, error) {
	file, err := os.OpenFile("apiserver_log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

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
	s.router.HandleFunc("/register", s.handleRegister()).Methods("POST")
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
			slog.Any("request_id", r.Context().Value(ctxKeyRequestId)))

		start := time.Now()
		rw := &responseWriter{ResponseWriter: w, code: http.StatusOK}

		next.ServeHTTP(rw, r)

		s.logger.Info(fmt.Sprintf("completed with %d %s in %v", rw.code, http.StatusText(rw.code), time.Since(start)))
	})
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
