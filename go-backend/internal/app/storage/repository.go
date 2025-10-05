package storage

import "github.com/go-next-pizza/internal/app/model"

type UserRepository interface {
	CreateUser(*model.User) (*model.User, error)
	FindByEmail(string) (*model.User, error)
	FindById(int) (*model.User, error)
}

type CartRepository interface {
	CreateCart(int, *model.Cart) (*model.Cart, error)
	AddProduct(int, *model.Cart) error 
	AddCombo(int, *model.Cart) error
	DeleteProduct(int, *model.Cart) error
	DeleteCombo(int, *model.Cart) error
	Refresh(*model.Cart) error
}

type OrderRepository interface {
	CreateOrder(*model.Cart,*model.Order) (*model.Order, error)
	AddProduct(int, int, *model.Order) error
	AddCombo(int, int, *model.Order) error
	GetOrderById(int) (*model.Order, error)
	GetOrdersByUserId(int) ([]*model.Order, error)
}

type CustomPizzaRepository interface {
	CreateCustomPizza(*model.CustomPizza) (*model.CustomPizza, error)
	GetCustomPizzaByID(int) (*model.CustomPizza, error)
	GetCustomPizzasByUserID(int) ([]*model.CustomPizza, error)
	UpdateCustomPizza(*model.CustomPizza) (*model.CustomPizza, error)
	DeleteCustomPizza(int) error
}

type ProductRepository interface {
	GetProductByID(int) (*model.Product, error)
	GetProducts() ([]*model.Product, error)
}

type IngredientRepository interface {
	GetIngredientByID(int) (*model.Ingredient, error)
	GetIngredients() ([]*model.Ingredient, error)
}