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
	
	AuthService        *service.AuthService
	CustomPizzaService *service.CustomPizzaService
	CartService        *service.CartService
	ProductService     *service.ProductService

	AuthHandler        *handler.AuthHandler
	CustomPizzaHandler *handler.CustomPizzaHandler
	ProductHandler 		 *handler.ProductHandler
	UserHandler 			 *handler.UserHandler
	OrderHandler 			 *handler.OrderHandler
	CartHandler 			 *handler.CartHandler
	ComboHandler 			 *handler.ComboHandler
	CategoryHandler 	 *handler.CategoryHandler	
	AuthMiddleware *middleware.AuthMiddleware
}

func New(cfg *config.Config) (*Container, error) {
	db, err := sql.Open("postgres", cfg.Database.URL)

	if err != nil {
		return nil, err
	}

	storage := sql_storage.NewSQLStorage(db)

	logger := slog.New(slog.NewJSONHandler(nil, nil))

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

	productService := service.NewProductService(storage.Product())

	cartService := service.NewCartService(
		storage.Cart(),
		storage.Product(),
		storage.Combo(),
	)

	comboService := service.NewComboService(storage.Combo(), storage.Product())

	authHandler := handler.NewAuthHandler(authService)
	customPizzaHandler := handler.NewCustomPizzaHandler(customPizzaService)
	productHandler := handler.NewProductHandler(productService)
	userHandler := handler.NewUserHandler(storage.User())
	orderHandler := handler.NewOrderHandler(storage.Order())
	cartHandler := handler.NewCartHandler(cartService)
	comboHandler := handler.NewComboHandler(comboService)
	categoryHandler := handler.NewCategoryHandler(storage.Category())
	authMiddleware := middleware.NewAuthMiddleware(authService)

	router := mux.NewRouter()
	
	// Добавляем CORS middleware для всех маршрутов
	router.Use(middleware.CORS)

	return &Container{
		Config:             cfg,
		DB:                 db,
		Storage:            storage,
		Logger:             logger,
		Router:             router,
		AuthService:        authService,
		CustomPizzaService: customPizzaService,
		CartService:        cartService,
		ProductService:     productService,
		AuthHandler:        authHandler,
		CustomPizzaHandler: customPizzaHandler,
		ProductHandler:     productHandler,
		UserHandler:				userHandler,
		OrderHandler:			  orderHandler,
		CartHandler:			  cartHandler,
		ComboHandler:			  comboHandler,
		CategoryHandler:  categoryHandler,
		AuthMiddleware:     authMiddleware,
	}, nil
}

func (c *Container) SetupRoutes() {
	c.Router.HandleFunc("/api/auth/register", c.AuthHandler.Register).Methods("POST")
	c.Router.HandleFunc("/api/auth/login", c.AuthHandler.Login).Methods("POST")

	c.Router.HandleFunc("/api/categories", c.CategoryHandler.GetCategories).Methods("GET")

	protected := c.Router.PathPrefix("/api").Subrouter()
	protected.Use(c.AuthMiddleware.RequireAuth)

	protected.HandleFunc("/cart/{id}", c.CartHandler.GetCart).Methods("GET")
	protected.HandleFunc("/cart/{id}/product", c.CartHandler.AddProduct).Methods("POST")
	protected.HandleFunc("/cart/{id}/combo", c.CartHandler.AddCombo).Methods("POST")
	protected.HandleFunc("/cart/{id}/product", c.CartHandler.DeleteProduct).Methods("DELETE")
	protected.HandleFunc("/cart/{id}/combo", c.CartHandler.DeleteCombo).Methods("DELETE")
	protected.HandleFunc("/cart/{id}/refresh", c.CartHandler.Refresh).Methods("PUT")

	protected.HandleFunc("/orders", c.OrderHandler.CreateOrder).Methods("POST")

	protected.HandleFunc("/users/{id}/orders", c.UserHandler.GetOrders).Methods("GET")

	protected.HandleFunc("/custom-pizzas", c.CustomPizzaHandler.CreateCustomPizza).Methods("POST")
	protected.HandleFunc("/custom-pizzas", c.CustomPizzaHandler.GetCustomPizzas).Methods("GET")
	protected.HandleFunc("/custom-pizzas/{id}", c.CustomPizzaHandler.GetCustomPizza).Methods("GET")
	protected.HandleFunc("/custom-pizzas/{id}", c.CustomPizzaHandler.UpdateCustomPizza).Methods("PUT")
	protected.HandleFunc("/custom-pizzas/{id}", c.CustomPizzaHandler.DeleteCustomPizza).Methods("DELETE")

	c.Router.HandleFunc("/api/products", c.ProductHandler.GetProducts).Methods("GET")
	c.Router.HandleFunc("/api/products/{id}", c.ProductHandler.GetProductById).Methods("GET")
	c.Router.HandleFunc("/api/products/category/{category}", c.ProductHandler.GetProductsByCategory).Methods("GET")

	c.Router.HandleFunc("/api/combos", c.ComboHandler.GetCombos).Methods("GET")
	c.Router.HandleFunc("/api/combos/{comboId}", c.ComboHandler.GetComboById).Methods("GET")
	c.Router.HandleFunc("/api/combos/{comboId}/replace", c.ComboHandler.ReplaceProduct).Methods("PUT")
}

func (c *Container) Start() error {
	c.SetupRoutes()
	
	server := &http.Server{
		Addr:    c.Config.Server.Host + ":" + c.Config.Server.Port,
		Handler: c.Router,
	}

	return server.ListenAndServe()
}