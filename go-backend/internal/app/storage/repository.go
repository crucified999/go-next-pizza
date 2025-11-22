package storage

import "github.com/go-next-pizza/internal/app/model"

type UserRepository interface {
	CreateUser(*model.User) (*model.User, error)
	FindByEmail(string) (*model.User, error)
	FindById(int) (*model.User, error)
	FindByPhone(string) (*model.User, error)
	GetOrders(int) ([]*model.Order, error)
	ChangeName(int, string) (error)
	ChangeEmail(int, string) (error)
}

type SMSCodeRepository interface {
    SaveCode(phone string, codeHash string, expiresAt int64) error
    GetLatestCodeHash(phone string) (codeHash string, expiresAt int64, err error)
    DeleteCodes(phone string) error
}

type CartRepository interface {
	CreateCart(int, *model.Cart) (*model.Cart, error)
	GetCartByUserId(int) (*model.Cart, error)
	GetCartProducts(int) ([]*model.CartProduct, error)
	GetCartCombos(int) ([]*model.CartCombo, error)
	AddProduct(int, int) error 
	AddCombo(int, int) error
	DeleteProduct(int, int) error
	DeleteCombo(int, int) error
	Refresh(int) error
}

type OrderRepository interface {
	CreateOrder(*model.Cart,*model.Order) (*model.Order, error)
	AddProduct(int, int, *model.Order) error
	AddCombo(int, int, *model.Order) error
	GetOrderById(int) (*model.Order, error)
	GetOrderProducts(int) ([]*model.Product, error)
	GetOrderCombos(int) ([]*model.Combo, error)
}

type CustomPizzaRepository interface {
	CreateCustomPizza(*model.CustomPizza) (*model.CustomPizza, error)
	UpdateDough(*model.CustomPizza,string) (*model.CustomPizza, error)
	UpdateSize(*model.CustomPizza,string) (*model.CustomPizza, error)
	GetCustomPizzaByID(int) (*model.CustomPizza, error)
	GetCustomPizzasByUserID(int) ([]*model.CustomPizza, error)
	UpdateCustomPizza(*model.CustomPizza) (*model.CustomPizza, error)
	DeleteCustomPizza(int) error
}

type ProductRepository interface {
	GetProducts() ([]*model.Product, error)
	GetProductById(int) (*model.Product, error)	
	GetProductsByCategory(int) ([]*model.Product, error)
	GetProductsVariants(int) ([]*model.ProductVariant, error)
	GetProductCategory(int) (string, error)
	ConvertIdToCategory(int) string
	ConvertCategorToId(string) int
	GetProductIngredients(int) ([]*model.Ingredient, error)
	GetProductToppings(int) ([]*model.Topping, error)
}

type ComboRepository interface {
	GetCombos() ([]*model.Combo, error)
	GetComboById(int) (*model.Combo, error)
	GetComboProducts(int) ([]*model.Product, error)
	GetComboDefaultProducts(int) ([]*model.Product, error)
	GetComboReplaces(int) (map[int][]int, error)
}

type IngredientRepository interface {
	GetIngredientByID(int) (*model.Ingredient, error)
	GetIngredients() ([]*model.Ingredient, error)
}

type CategoryRepository interface {
	GetCategories() ([]*model.Category, error)
}