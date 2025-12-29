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
	SMSAuthService     *service.SMSAuthService
	CustomPizzaService *service.CustomPizzaService
	CartService        *service.CartService
	ProductService     *service.ProductService

	AuthHandler        *handler.AuthHandler
	SMSAuthHandler     *handler.SMSAuthHandler
	CustomPizzaHandler *handler.CustomPizzaHandler
	ProductHandler 		 *handler.ProductHandler
	UserHandler 			 *handler.UserHandler
	OrderHandler 			 *handler.OrderHandler
	CartHandler 			 *handler.CartHandler
	ComboHandler 			 *handler.ComboHandler
	CategoryHandler 	 *handler.CategoryHandler	
	AuthMiddleware 			*middleware.AuthMiddleware
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
		storage.Cart(),
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

  smsAuthService := service.NewSMSAuthServiceWithCodes(storage.User(), storage.SMSCode())

	cartService := service.NewCartService(
		storage.Cart(),
		storage.Product(),
		storage.Combo(),
	)

	comboService := service.NewComboService(storage.Combo(), storage.Product())

	authHandler := handler.NewAuthHandler(authService, storage.User())
  smsAuthHandler := handler.NewSMSAuthHandler(smsAuthService, authService)
	customPizzaHandler := handler.NewCustomPizzaHandler(customPizzaService)
	productHandler := handler.NewProductHandler(productService)
	userHandler := handler.NewUserHandler(storage.User())
	orderHandler := handler.NewOrderHandler(storage.Order())
	cartHandler := handler.NewCartHandler(cartService, productService)
	comboHandler := handler.NewComboHandler(comboService)
	categoryHandler := handler.NewCategoryHandler(storage.Category())
	authMiddleware := middleware.NewAuthMiddleware(authService)

	router := mux.NewRouter()
	
	router.Use(middleware.CORS)
	router.Use(middleware.CORSOptions)

	return &Container{
		Config:             cfg,
		DB:                 db,
		Storage:            storage,
		Logger:             logger,
		Router:             router,
		AuthService:        authService,
		SMSAuthService:     smsAuthService,
		CustomPizzaService: customPizzaService,
		CartService:        cartService,
		ProductService:     productService,
		AuthHandler:        authHandler,
		SMSAuthHandler:     smsAuthHandler,
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

func optionsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func (c *Container) SetupRoutes() {
	c.Router.HandleFunc("/api/auth/register", optionsHandler).Methods("OPTIONS")
  c.Router.HandleFunc("/api/auth/login", optionsHandler).Methods("OPTIONS")
  c.Router.HandleFunc("/api/auth/check", optionsHandler).Methods("OPTIONS")
  c.Router.HandleFunc("/api/categories", optionsHandler).Methods("OPTIONS")
  c.Router.HandleFunc("/api/auth/sms/send", optionsHandler).Methods("OPTIONS")
  c.Router.HandleFunc("/api/auth/sms/verify", optionsHandler).Methods("OPTIONS")
  c.Router.HandleFunc("/api/products", optionsHandler).Methods("OPTIONS")
  c.Router.HandleFunc("/api/products/{id}", optionsHandler).Methods("OPTIONS")
  c.Router.HandleFunc("/api/products/category/{category}", optionsHandler).Methods("OPTIONS")
	c.Router.HandleFunc("/api/products/variant", optionsHandler).Methods("OPTIONS")
  c.Router.HandleFunc("/api/combos", optionsHandler).Methods("OPTIONS")
  c.Router.HandleFunc("/api/combos/{comboId}", optionsHandler).Methods("OPTIONS")
  c.Router.HandleFunc("/api/combos/{comboId}/replace", optionsHandler).Methods("OPTIONS")

	c.Router.HandleFunc("/api/auth/register", c.AuthHandler.Register).Methods("POST")
	c.Router.HandleFunc("/api/auth/login", c.AuthHandler.Login).Methods("POST")
	c.Router.HandleFunc("/api/auth/check", c.AuthHandler.CheckAuth).Methods("GET")
	c.Router.HandleFunc("/api/auth/refresh", c.AuthHandler.RefreshToken).Methods("POST")

	c.Router.HandleFunc("/api/categories", c.CategoryHandler.GetCategories).Methods("GET")

	protected := c.Router.PathPrefix("/api").Subrouter()
	protected.Use(c.AuthMiddleware.RequireAuth)

	protected.HandleFunc("/cart", optionsHandler).Methods("OPTIONS")
  protected.HandleFunc("/cart/add-product", optionsHandler).Methods("OPTIONS")
	protected.HandleFunc("/cart/add-pizza", optionsHandler).Methods("OPTIONS")
  protected.HandleFunc("/cart/{id}/combo", optionsHandler).Methods("OPTIONS")
  protected.HandleFunc("/cart/delete-product", optionsHandler).Methods("OPTIONS")
	protected.HandleFunc("/cart/delete-pizza", optionsHandler).Methods("OPTIONS")
  protected.HandleFunc("/cart/{id}/combo", optionsHandler).Methods("OPTIONS")
  protected.HandleFunc("/cart/refresh", optionsHandler).Methods("OPTIONS")
  protected.HandleFunc("/auth/logout", optionsHandler).Methods("OPTIONS")
  protected.HandleFunc("/orders", optionsHandler).Methods("OPTIONS")
  protected.HandleFunc("/users/orders", optionsHandler).Methods("OPTIONS")
	protected.HandleFunc("/users/{id}/name", optionsHandler).Methods("OPTIONS")
	protected.HandleFunc("/users/{id}/email", optionsHandler).Methods("OPTIONS")
  protected.HandleFunc("/custom-pizzas", optionsHandler).Methods("OPTIONS")
  protected.HandleFunc("/custom-pizzas/{id}", optionsHandler).Methods("OPTIONS")
	protected.HandleFunc("/custom-pizzas/{id}/dough", optionsHandler).Methods("OPTIONS")

	protected.HandleFunc("/cart", c.CartHandler.GetCart).Methods("GET")
	protected.HandleFunc("/cart/add-product", c.CartHandler.AddProduct).Methods("POST")
	protected.HandleFunc("/cart/add-pizza", c.CartHandler.AddPizza).Methods("POST")
	protected.HandleFunc("/cart/{id}/combo", c.CartHandler.AddCombo).Methods("POST")
	protected.HandleFunc("/cart/delete-product", c.CartHandler.DeleteProduct).Methods("DELETE")
	protected.HandleFunc("/cart/delete-pizza", c.CartHandler.DeletePizza).Methods("DELETE")
	protected.HandleFunc("/cart/{id}/combo", c.CartHandler.DeleteCombo).Methods("DELETE")
	protected.HandleFunc("/cart/refresh", c.CartHandler.Refresh).Methods("PUT")

	protected.HandleFunc("/orders", c.OrderHandler.CreateOrder).Methods("POST")

	protected.HandleFunc("/users/orders", c.UserHandler.GetOrders).Methods("GET")
	protected.HandleFunc("/users/{id}/name", c.UserHandler.ChangeName).Methods("PATCH")
	protected.HandleFunc("/users/{id}/email", c.UserHandler.ChangeEmail).Methods("PATCH")

	protected.HandleFunc("/custom-pizzas", c.CustomPizzaHandler.CreateCustomPizza).Methods("POST")
	protected.HandleFunc("/custom-pizzas", c.CustomPizzaHandler.GetCustomPizzas).Methods("GET")
	protected.HandleFunc("/custom-pizzas/{id}", c.CustomPizzaHandler.GetCustomPizza).Methods("GET")
	protected.HandleFunc("/custom-pizzas/{id}", c.CustomPizzaHandler.UpdateCustomPizza).Methods("PUT")
	protected.HandleFunc("/custom-pizzas/{id}/dough", c.CustomPizzaHandler.SetDough).Methods("PUT")
	protected.HandleFunc("/custom-pizzas/{id}", c.CustomPizzaHandler.DeleteCustomPizza).Methods("DELETE")

	c.Router.HandleFunc("/api/products", c.ProductHandler.GetProducts).Methods("GET")
	c.Router.HandleFunc("/api/products/variant", c.ProductHandler.GetProductVariant).Methods("GET")
	c.Router.HandleFunc("/api/products/{id}", c.ProductHandler.GetProductById).Methods("GET")
	c.Router.HandleFunc("/api/products/category/{category}", c.ProductHandler.GetProductsByCategory).Methods("GET")

	c.Router.HandleFunc("/api/combos", c.ComboHandler.GetCombos).Methods("GET")
	c.Router.HandleFunc("/api/combos/{comboId}", c.ComboHandler.GetComboById).Methods("GET")
	c.Router.HandleFunc("/api/combos/{comboId}/replace", c.ComboHandler.ReplaceProduct).Methods("PUT")

	c.Router.HandleFunc("/api/auth/sms/send", c.SMSAuthHandler.SendCode).Methods("POST")
	c.Router.HandleFunc("/api/auth/sms/verify", c.SMSAuthHandler.VerifyCode).Methods("POST")

	protected.HandleFunc("/auth/logout", c.AuthHandler.Logout).Methods("POST")
}

func (c *Container) Start() error {
	c.SetupRoutes()
	
	server := &http.Server{
		Addr:    c.Config.Server.Host + ":" + c.Config.Server.Port,
		Handler: c.Router,
	}

	return server.ListenAndServe()
}