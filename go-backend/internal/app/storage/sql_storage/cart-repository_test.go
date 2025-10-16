package sql_storage_test

import (
	"testing"

	"github.com/go-next-pizza/internal/app/model"
	sqlstorage "github.com/go-next-pizza/internal/app/storage/sql_storage"
	"github.com/stretchr/testify/assert"
)

func TestCartRepository_CreateCart(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("carts", "users")

	s := sqlstorage.NewSQLStorage(db)
	u := model.TestUser(t)
	c := model.TestCart(t)

	user, err := s.User().CreateUser(u)
	assert.NoError(t, err)

	cart, err := s.Cart().CreateCart(user.Id, c)

	assert.NoError(t, err)
	assert.NotNil(t, cart)
	assert.Equal(t, c.Id, cart.Id)
}

func TestCartRepository_AddProduct(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("carts", "products_in_cart", "products", "users")

	s := sqlstorage.NewSQLStorage(db)
	u := model.TestUser(t)
	c := model.TestCart(t)

	user, err := s.User().CreateUser(u)
	assert.NoError(t, err)

	cart, err := s.Cart().CreateCart(user.Id, c)
	assert.NoError(t, err)
	assert.Equal(t, cart.Id, c.Id)

	var productId int
	
	err = db.QueryRow("INSERT INTO products (title, price) VALUES ($1, $2) RETURNING id", "Test Product", 100).Scan(&productId)
	assert.NoError(t, err)

	err = s.Cart().AddProduct(productId, c.Id)
	assert.NoError(t, err)

	err = s.Cart().AddProduct(productId, c.Id)
	assert.NoError(t, err)

	var cnt int

	err = db.QueryRow("SELECT amount FROM products_in_cart WHERE cart_id = $1 AND product_id = $2", c.Id, productId).Scan(&cnt)
	assert.NoError(t, err)
	assert.Equal(t, 2, cnt)
}

func TestCartRepository_AddCombo(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("carts", "combos_in_cart", "combos", "users")

	s := sqlstorage.NewSQLStorage(db)
	u := model.TestUser(t)
	c := model.TestCart(t)

	user, err := s.User().CreateUser(u)
	assert.NoError(t, err)

	cart, err := s.Cart().CreateCart(user.Id, c)
	assert.NoError(t, err)
	assert.Equal(t, cart.Id, c.Id)

	var comboId int
	err = db.QueryRow("INSERT INTO combos (title, description) VALUES ($1, $2) RETURNING id", "Test Combo", "Test Description").Scan(&comboId)
	assert.NoError(t, err)

	err = s.Cart().AddCombo(comboId, c.Id)
	assert.NoError(t, err)

	err = s.Cart().AddCombo(comboId, c.Id)
	assert.NoError(t, err)

	var cnt int

	err = db.QueryRow("SELECT amount FROM combos_in_cart WHERE cart_id = $1 AND combo_id = $2", c.Id, comboId).Scan(&cnt)
	assert.NoError(t, err)
	assert.Equal(t, 2, cnt)
}

func TestCartRepository_DeleteProduct(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("carts", "products_in_cart", "products", "users")

	s := sqlstorage.NewSQLStorage(db)
	u := model.TestUser(t)
	c := model.TestCart(t)

	user, err := s.User().CreateUser(u)
	assert.NoError(t, err)

	cart, err := s.Cart().CreateCart(user.Id, c)
	assert.NoError(t, err)
	assert.Equal(t, cart.Id, c.Id)

	var productId int
	err = db.QueryRow("INSERT INTO products (title, price) VALUES ($1, $2) RETURNING id", "Test Product", 100).Scan(&productId)
	assert.NoError(t, err)

	err = s.Cart().AddProduct(productId, c.Id)
	assert.NoError(t, err)

	err = s.Cart().DeleteProduct(productId, c.Id)
	assert.NoError(t, err)

	var cnt int

	err = db.QueryRow("SELECT amount FROM products_in_cart WHERE cart_id = $1 AND product_id = $2", c.Id, productId).Scan(&cnt)
	assert.Error(t, err)
	assert.Equal(t, 0, cnt)

	err = s.Cart().AddProduct(productId, c.Id)
	assert.NoError(t, err)

	err = s.Cart().AddProduct(productId, c.Id)
	assert.NoError(t, err)

	err = s.Cart().DeleteProduct(productId, c.Id)
	assert.NoError(t, err)

	err = db.QueryRow("SELECT amount FROM products_in_cart WHERE cart_id = $1 AND product_id = $2", c.Id, productId).Scan(&cnt)
	assert.NoError(t, err)
	assert.Equal(t, 1, cnt)
}

func TestCartRepository_DeleteCombo(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("carts", "products_in_cart", "products", "users")

	s := sqlstorage.NewSQLStorage(db)
	u := model.TestUser(t)
	c := model.TestCart(t)

	user, err := s.User().CreateUser(u)
	assert.NoError(t, err)

	cart, err := s.Cart().CreateCart(user.Id, c)
	assert.NoError(t, err)
	assert.Equal(t, cart.Id, c.Id)

	var comboId int
	err = db.QueryRow("INSERT INTO combos (title, description) VALUES ($1, $2) RETURNING id", "Test Combo", "Test Description").Scan(&comboId)
	assert.NoError(t, err)

	err = s.Cart().AddCombo(comboId, c.Id)
	assert.NoError(t, err)

	err = s.Cart().DeleteCombo(comboId, c.Id)
	assert.NoError(t, err)

	var cnt int

	err = db.QueryRow("SELECT amount FROM combos_in_cart WHERE cart_id = $1 AND combo_id = $2", c.Id, comboId).Scan(&cnt)
	assert.Error(t, err)
	assert.Equal(t, 0, cnt)

	err = s.Cart().AddCombo(comboId, c.Id)
	assert.NoError(t, err)

	err = s.Cart().AddCombo(comboId, c.Id)
	assert.NoError(t, err)

	err = s.Cart().DeleteCombo(comboId, c.Id)
	assert.NoError(t, err)

	err = db.QueryRow("SELECT amount FROM combos_in_cart WHERE cart_id = $1 AND combo_id = $2", c.Id, comboId).Scan(&cnt)
	assert.NoError(t, err)
	assert.Equal(t, 1, cnt)
}

func TestCartRepository_Refresh(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("carts", "products_in_cart", "products", "users")

	s := sqlstorage.NewSQLStorage(db)
	u := model.TestUser(t)
	c := model.TestCart(t)

	user, err := s.User().CreateUser(u)
	assert.NoError(t, err)

	cart, err := s.Cart().CreateCart(user.Id, c)
	assert.NoError(t, err)
	assert.Equal(t, cart.Id, c.Id)

	var productId int
	var comboId int
	err = db.QueryRow("INSERT INTO products (title, price) VALUES ($1, $2) RETURNING id", "Test Product", 100).Scan(&productId)
	assert.NoError(t, err)

	err = db.QueryRow("INSERT INTO combos (title, description) VALUES ($1, $2) RETURNING id", "Test Combo", "Test Description").Scan(&comboId)
	assert.NoError(t, err)

	err = s.Cart().AddProduct(productId, c.Id)
	assert.NoError(t, err)

	err = s.Cart().AddCombo(comboId, c.Id)
	assert.NoError(t, err)

	err = s.Cart().Refresh(c.Id)
	assert.NoError(t, err)

	var productCount int
	var comboCount int

	err = db.QueryRow("SELECT amount FROM products_in_cart WHERE cart_id = $1 AND product_id = $2", c.Id, productId).Scan(&productCount)
	assert.Error(t, err)
	assert.Equal(t, 0, productCount)

	err = db.QueryRow("SELECT amount FROM combos_in_cart WHERE cart_id = $1 AND combo_id = $2", c.Id, comboId).Scan(&comboCount)
	assert.Error(t, err)
	assert.Equal(t, 0, comboCount)
}