package sql_storage_test

import (
	"testing"

	"github.com/go-next-pizza/internal/app/model"
	sqlstorage "github.com/go-next-pizza/internal/app/storage/sql_storage"
	"github.com/stretchr/testify/assert"
)

func TestOrderRepository_CreateOrder(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("orders", "products_in_order", "products", "users", "carts", "products_in_cart")

	s := sqlstorage.NewSQLStorage(db)

	u := model.TestUser(t)
	c := &model.Cart{}
	createdUser, err := s.User().CreateUser(u)
	assert.NoError(t, err)

	cart, err := s.Cart().CreateCart(createdUser.Id, c)
	assert.NoError(t, err)

	var productId int
	err = db.QueryRow("INSERT INTO products (title, price) VALUES ($1, $2) RETURNING id", "Test Product", 100).Scan(&productId)
	assert.NoError(t, err)

	var comboId int
	err = db.QueryRow("INSERT INTO combos (title, description) VALUES ($1, $2) RETURNING id", "Test Combo", "Test Description").Scan(&comboId)
	assert.NoError(t, err)

	err = s.Cart().AddProduct(productId, "1 шт", cart.Id)
	assert.NoError(t, err)

	err = s.Cart().AddProduct(productId, "1 шт", cart.Id)
	assert.NoError(t, err)

	err = s.Cart().AddCombo(comboId, cart.Id)
	assert.NoError(t, err)

	order := model.TestOrder(createdUser.Id, t)

	err = s.Order().CreateOrder(order)

	assert.NoError(t, err)

	var cnt int
	
	err = db.QueryRow("SELECT COUNT(*) FROM orders WHERE id = $1", order.Id).Scan(&cnt)
	assert.NoError(t, err)
	assert.Equal(t, 1, cnt)

	err = db.QueryRow("SELECT count FROM products_in_order WHERE order_id = $1 AND product_id = $2 AND amount = $3", order.Id, productId, "1 шт").Scan(&cnt)

	assert.NoError(t, err)
	assert.Equal(t, 2, cnt)

	err = db.QueryRow("SELECT amount FROM combos_in_order WHERE order_id = $1 AND combo_id = $2", order.Id, comboId).Scan(&cnt)

	assert.NoError(t, err)
	assert.Equal(t, 1, cnt)
}

// func TestOrderRepository_AddProduct(t *testing.T) {
// 	db, teardown := sqlstorage.TestDB(t, databaseURL)
// 	defer teardown("orders", "products_in_order")

// 	s := sqlstorage.NewSQLStorage(db)

// 	user := model.TestUser(t)
// 	assert.NoError(t, s.User().CreateUser(user))

// 	order := model.TestOrder(user.Id, t)
// 	assert.NoError(t, s.Order().CreateOrder(order))

// 	product := model.TestProduct(t)
// 	assert.NoError(t, s.Product().CreateProduct(product))
	
	
// }