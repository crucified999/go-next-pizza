package sql_storage_test

import (
	"testing"

	"github.com/go-next-pizza/internal/app/model"
	sqlstorage "github.com/go-next-pizza/internal/app/storage/sql_storage"
	"github.com/stretchr/testify/assert"
)

func TestCustomPizzaRepository_CreateCustomPizza(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("custom_pizzas", "custom_pizza_ingredients", "users", "products", "pizza_ingredients")

	s := sqlstorage.NewSQLStorage(db)

	// Создаем пользователя
	u := model.TestUser(t)
	user, err := s.User().CreateUser(u)
	assert.NoError(t, err)

	// Создаем продукт (базовую пиццу)
	var productId int
	err = db.QueryRow("INSERT INTO products (title, price) VALUES ($1, $2) RETURNING id", "Test Pizza", 100.0).Scan(&productId)
	assert.NoError(t, err)

	// Создаем ингредиент
	var ingredientId int
	err = db.QueryRow("INSERT INTO pizza_ingredients (title, product_id) VALUES ($1, $2) RETURNING id", "Test Ingredient", productId).Scan(&ingredientId)
	assert.NoError(t, err)

	// Создаем кастомную пиццу
	customPizza := &model.CustomPizza{
		UserID:      user.Id,
		BasePizzaID: productId,
		Name:        "My Custom Pizza",
		TotalPrice:  120.0,
		Ingredients: []model.CustomPizzaIngredient{
			{
				IngredientID: ingredientId,
				IsAdded:      true,
			},
		},
	}

	createdCustomPizza, err := s.CustomPizza().CreateCustomPizza(customPizza)
	assert.NoError(t, err)
	assert.NotNil(t, createdCustomPizza)
	assert.Equal(t, customPizza.UserID, createdCustomPizza.UserID)
	assert.Equal(t, customPizza.BasePizzaID, createdCustomPizza.BasePizzaID)
	assert.Equal(t, customPizza.Name, createdCustomPizza.Name)
	assert.Equal(t, customPizza.TotalPrice, createdCustomPizza.TotalPrice)
	assert.Len(t, createdCustomPizza.Ingredients, 1)
}

func TestCustomPizzaRepository_GetCustomPizzaByID(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("custom_pizzas", "custom_pizza_ingredients", "users", "products", "pizza_ingredients")

	s := sqlstorage.NewSQLStorage(db)

	// Создаем пользователя
	u := model.TestUser(t)
	user, err := s.User().CreateUser(u)
	assert.NoError(t, err)

	// Создаем продукт
	var productId int
	err = db.QueryRow("INSERT INTO products (title, price) VALUES ($1, $2) RETURNING id", "Test Pizza", 100.0).Scan(&productId)
	assert.NoError(t, err)

	// Создаем кастомную пиццу
	var customPizzaId int
	err = db.QueryRow("INSERT INTO custom_pizzas (user_id, base_pizza_id, name, total_price) VALUES ($1, $2, $3, $4) RETURNING id", user.Id, productId, "Test Custom Pizza", 120.0).Scan(&customPizzaId)
	assert.NoError(t, err)

	// Получаем кастомную пиццу
	customPizza, err := s.CustomPizza().GetCustomPizzaByID(customPizzaId)
	assert.NoError(t, err)
	assert.NotNil(t, customPizza)
	assert.Equal(t, customPizzaId, customPizza.ID)
	assert.Equal(t, user.Id, customPizza.UserID)
	assert.Equal(t, productId, customPizza.BasePizzaID)
	assert.Equal(t, "Test Custom Pizza", customPizza.Name)
	assert.Equal(t, 120.0, customPizza.TotalPrice)
}

func TestCustomPizzaRepository_GetCustomPizzasByUserID(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("custom_pizzas", "custom_pizza_ingredients", "users", "products", "pizza_ingredients")

	s := sqlstorage.NewSQLStorage(db)

	// Создаем пользователя
	u := model.TestUser(t)
	user, err := s.User().CreateUser(u)
	assert.NoError(t, err)

	// Создаем продукт
	var productId int
	err = db.QueryRow("INSERT INTO products (title, price) VALUES ($1, $2) RETURNING id", "Test Pizza", 100.0).Scan(&productId)
	assert.NoError(t, err)

	// Создаем несколько кастомных пицц
	_, err = db.Exec("INSERT INTO custom_pizzas (user_id, base_pizza_id, name, total_price) VALUES ($1, $2, $3, $4)", user.Id, productId, "Custom Pizza 1", 120.0)
	assert.NoError(t, err)

	_, err = db.Exec("INSERT INTO custom_pizzas (user_id, base_pizza_id, name, total_price) VALUES ($1, $2, $3, $4)", user.Id, productId, "Custom Pizza 2", 130.0)
	assert.NoError(t, err)

	// Получаем кастомные пиццы пользователя
	customPizzas, err := s.CustomPizza().GetCustomPizzasByUserID(user.Id)
	assert.NoError(t, err)
	assert.Len(t, customPizzas, 2)
}

func TestCustomPizzaRepository_UpdateCustomPizza(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("custom_pizzas", "custom_pizza_ingredients", "users", "products", "pizza_ingredients")

	s := sqlstorage.NewSQLStorage(db)

	// Создаем пользователя
	u := model.TestUser(t)
	user, err := s.User().CreateUser(u)
	assert.NoError(t, err)

	// Создаем продукт
	var productId int
	err = db.QueryRow("INSERT INTO products (title, price) VALUES ($1, $2) RETURNING id", "Test Pizza", 100.0).Scan(&productId)
	assert.NoError(t, err)

	// Создаем кастомную пиццу
	customPizza := &model.CustomPizza{
		UserID:      user.Id,
		BasePizzaID: productId,
		Name:        "Original Name",
		TotalPrice:  120.0,
		Ingredients: []model.CustomPizzaIngredient{},
	}

	createdCustomPizza, err := s.CustomPizza().CreateCustomPizza(customPizza)
	assert.NoError(t, err)

	// Обновляем кастомную пиццу
	createdCustomPizza.Name = "Updated Name"
	createdCustomPizza.TotalPrice = 140.0

	updatedCustomPizza, err := s.CustomPizza().UpdateCustomPizza(createdCustomPizza)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Name", updatedCustomPizza.Name)
	assert.Equal(t, 140.0, updatedCustomPizza.TotalPrice)
}

func TestCustomPizzaRepository_DeleteCustomPizza(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("custom_pizzas", "custom_pizza_ingredients", "users", "products", "pizza_ingredients")

	s := sqlstorage.NewSQLStorage(db)

	// Создаем пользователя
	u := model.TestUser(t)
	user, err := s.User().CreateUser(u)
	assert.NoError(t, err)

	// Создаем продукт
	var productId int
	err = db.QueryRow("INSERT INTO products (title, price) VALUES ($1, $2) RETURNING id", "Test Pizza", 100.0).Scan(&productId)
	assert.NoError(t, err)

	// Создаем кастомную пиццу
	var customPizzaId int
	err = db.QueryRow("INSERT INTO custom_pizzas (user_id, base_pizza_id, name, total_price) VALUES ($1, $2, $3, $4) RETURNING id", user.Id, productId, "Test Custom Pizza", 120.0).Scan(&customPizzaId)
	assert.NoError(t, err)

	// Удаляем кастомную пиццу
	err = s.CustomPizza().DeleteCustomPizza(customPizzaId)
	assert.NoError(t, err)

	// Проверяем, что кастомная пицца удалена
	_, err = s.CustomPizza().GetCustomPizzaByID(customPizzaId)
	assert.Error(t, err)
}