package sql_storage

import (
	"database/sql"

	"github.com/go-next-pizza/internal/app/storage"
	_ "github.com/lib/pq"
)

type SQLStorage struct {
	db *sql.DB
	userRepository *UserRepository
	cartRepository *CartRepository
	orderRepository *OrderRepository
	customPizzaRepository *CustomPizzaRepository
	productRepository *ProductRepository
	ingredientRepository *IngredientRepository
}

func NewSQLStorage(db *sql.DB) *SQLStorage {
	return &SQLStorage{
		db: db,
	}
}

func (s *SQLStorage) User() storage.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		storage: s,
	}

	return s.userRepository
}

func (s *SQLStorage) Cart() storage.CartRepository {
	if s.cartRepository != nil {
		return s.cartRepository
	}

	s.cartRepository = &CartRepository{
		storage: s,
	}

	return s.cartRepository
}

func (s *SQLStorage) Order() storage.OrderRepository {
	if s.orderRepository != nil {
		return s.orderRepository
	}

	s.orderRepository = &OrderRepository{
		storage: s,
	}

	return s.orderRepository
}

func (s *SQLStorage) CustomPizza() storage.CustomPizzaRepository {
	if s.customPizzaRepository != nil {
		return s.customPizzaRepository
	}

	s.customPizzaRepository = &CustomPizzaRepository{
		storage: s,
	}

	return s.customPizzaRepository
}

func (s *SQLStorage) Product() storage.ProductRepository {
	if s.productRepository != nil {
		return s.productRepository
	}

	s.productRepository = &ProductRepository{
		storage: s,
	}

	return s.productRepository
}

func (s *SQLStorage) Ingredient() storage.IngredientRepository {
	if s.ingredientRepository != nil {
		return s.ingredientRepository
	}

	s.ingredientRepository = &IngredientRepository{
		storage: s,
	}

	return s.ingredientRepository
}