package container

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/go-next-pizza/internal/app/config"
	"github.com/go-next-pizza/internal/app/handler"
	"github.com/go-next-pizza/internal/app/middleware"
	"github.com/go-next-pizza/internal/app/service"
	"github.com/go-next-pizza/internal/app/storage/sql_storage"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Container struct {
	Config      *config.Config
	DB          *sql.DB
	Storage     *sql_storage.SQLStorage
	Logger      *slog.Logger
	Router      *mux.Router
	
	// Services
	AuthService        *service.AuthService
	CustomPizzaService *service.CustomPizzaService
	
	// Handlers
	AuthHandler        *handler.AuthHandler
	CustomPizzaHandler *handler.CustomPizzaHandler
	
	// Middleware
	AuthMiddleware *middleware.AuthMiddleware
}

func New(cfg *config.Config) (*Container, error) {
	// Database
	db, err := sql.Open("postgres", cfg.Database.URL)
	if err != nil {
		return nil, err
	}

	// Storage
	storage := sql_storage.NewSQLStorage(db)

	// Logger
	logger := slog.New(slog.NewJSONHandler(nil, nil))

	// Services
	authService := service.NewAuthService(
		storage.User(),
		[]byte(cfg.JWT.Secret),
		cfg.JWT.TTLSeconds,
		cfg.JWT.RefreshTTLSeconds,
	)

	customPizzaService := service.NewCustomPizzaService(
		storage.CustomPizza(),
		storage.Product(),
		storage.Ingredient(),
	)

	// Handlers
	authHandler := handler.NewAuthHandler(authService)
	customPizzaHandler := handler.NewCustomPizzaHandler(customPizzaService)

	// Middleware
	authMiddleware := middleware.NewAuthMiddleware(authService)

	// Router
	router := mux.NewRouter()

	return &Container{
		Config:             cfg,
		DB:                 db,
		Storage:            storage,
		Logger:             logger,
		Router:             router,
		AuthService:        authService,
		CustomPizzaService: customPizzaService,
		AuthHandler:        authHandler,
		CustomPizzaHandler: customPizzaHandler,
		AuthMiddleware:     authMiddleware,
	}, nil
}

func (c *Container) SetupRoutes() {
	// Public routes
	c.Router.HandleFunc("/api/auth/register", c.AuthHandler.Register).Methods("POST")
	c.Router.HandleFunc("/api/auth/login", c.AuthHandler.Login).Methods("POST")

	// Protected routes
	protected := c.Router.PathPrefix("/api").Subrouter()
	protected.Use(c.AuthMiddleware.RequireAuth)

	protected.HandleFunc("/custom-pizzas", c.CustomPizzaHandler.CreateCustomPizza).Methods("POST")
	protected.HandleFunc("/custom-pizzas", c.CustomPizzaHandler.GetCustomPizzas).Methods("GET")
	protected.HandleFunc("/custom-pizzas/{id}", c.CustomPizzaHandler.GetCustomPizza).Methods("GET")
	protected.HandleFunc("/custom-pizzas/{id}", c.CustomPizzaHandler.UpdateCustomPizza).Methods("PUT")
	protected.HandleFunc("/custom-pizzas/{id}", c.CustomPizzaHandler.DeleteCustomPizza).Methods("DELETE")
}

func (c *Container) Start() error {
	c.SetupRoutes()
	
	server := &http.Server{
		Addr:    c.Config.Server.Host + ":" + c.Config.Server.Port,
		Handler: c.Router,
	}

	return server.ListenAndServe()
}