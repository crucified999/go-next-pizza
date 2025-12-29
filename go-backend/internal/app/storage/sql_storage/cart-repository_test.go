package sql_storage_test

import (
	"database/sql"
	"testing"

	"github.com/go-next-pizza/internal/app/model"
	sqlstorage "github.com/go-next-pizza/internal/app/storage/sql_storage"
	"github.com/stretchr/testify/assert"
)

func TestCartRepository_CreateCart(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("carts", "users")

	s := sqlstorage.NewSQLStorage(db)
	
	userID, err := createTestUser(db)
	assert.NoError(t, err)
	
	cart := &model.Cart{
		UserId: userID,
	}
	
	createdCart, err := s.Cart().CreateCart(userID, cart)
	
	assert.NoError(t, err)
	assert.NotNil(t, createdCart)
	assert.Equal(t, userID, createdCart.UserId)
	assert.NotZero(t, createdCart.Id)
}

func TestCartRepository_AddProduct(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("carts", "products_in_cart", "products", "users")

	s := sqlstorage.NewSQLStorage(db)
	
	userID, err := createTestUser(db)
	assert.NoError(t, err)
	
	cart := &model.Cart{UserId: userID}
	cart, err = s.Cart().CreateCart(userID, cart)
	assert.NoError(t, err)
	
	var productId int
	err = db.QueryRow(
		"INSERT INTO products (title, price) VALUES ($1, $2) RETURNING id", 
		"Test Product", 
		100,
	).Scan(&productId)
	assert.NoError(t, err)
	
	err = s.Cart().AddProduct(productId, "1 шт", cart.Id)
	assert.NoError(t, err)
	
	err = s.Cart().AddProduct(productId, "1 шт", cart.Id)
	assert.NoError(t, err)
	
	var count int
	err = db.QueryRow(
		"SELECT count FROM products_in_cart WHERE cart_id = $1 AND product_id = $2 AND amount = $3", 
		cart.Id, 
		productId, 
		"1 шт",
	).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 2, count)
}

func TestCartRepository_DeleteProduct(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("carts", "products_in_cart", "products", "users")

	s := sqlstorage.NewSQLStorage(db)
	
	userID, err := createTestUser(db)
	assert.NoError(t, err)
	
	cart := &model.Cart{UserId: userID}
	cart, err = s.Cart().CreateCart(userID, cart)
	assert.NoError(t, err)
	
	var productId int
	err = db.QueryRow(
		"INSERT INTO products (title, price) VALUES ($1, $2) RETURNING id", 
		"Test Product", 
		100,
	).Scan(&productId)
	assert.NoError(t, err)
	
	err = s.Cart().AddProduct(productId, "1 шт", cart.Id)
	assert.NoError(t, err)
	
	err = s.Cart().DeleteProduct(productId, cart.Id, "1 шт")
	assert.NoError(t, err)
	
	var count int
	err = db.QueryRow(
		"SELECT count FROM products_in_cart WHERE cart_id = $1 AND product_id = $2", 
		cart.Id, 
		productId,
	).Scan(&count)
	assert.Error(t, err) 
	
	err = s.Cart().AddProduct(productId, "1 шт", cart.Id)
	assert.NoError(t, err)
	err = s.Cart().AddProduct(productId, "1 шт", cart.Id)
	assert.NoError(t, err)
	
	err = s.Cart().DeleteProduct(productId, cart.Id, "1 шт")
	assert.NoError(t, err)
	
	err = db.QueryRow(
		"SELECT count FROM products_in_cart WHERE cart_id = $1 AND product_id = $2 AND amount = $3", 
		cart.Id, 
		productId, 
		"1 шт",
	).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)
}

func TestCartRepository_AddCombo(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("carts", "combos_in_cart", "combos", "users")

	s := sqlstorage.NewSQLStorage(db)
	
	userID, err := createTestUser(db)
	assert.NoError(t, err)
	
	cart := &model.Cart{UserId: userID}
	cart, err = s.Cart().CreateCart(userID, cart)
	assert.NoError(t, err)
	
	var comboId int
	err = db.QueryRow(
		"INSERT INTO combos (title, price) VALUES ($1, $2) RETURNING id", 
		"Test Combo", 
		500,
	).Scan(&comboId)
	assert.NoError(t, err)
	
	err = s.Cart().AddCombo(comboId, cart.Id)
	assert.NoError(t, err)
	
	err = s.Cart().AddCombo(comboId, cart.Id)
	assert.NoError(t, err)
	
	var amount int
	err = db.QueryRow(
		"SELECT amount FROM combos_in_cart WHERE cart_id = $1 AND combo_id = $2", 
		cart.Id, 
		comboId,
	).Scan(&amount)
	assert.NoError(t, err)
	assert.Equal(t, 2, amount)
}

func TestCartRepository_DeleteCombo(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("carts", "combos_in_cart", "combos", "users")

	s := sqlstorage.NewSQLStorage(db)
	
	userID, err := createTestUser(db)
	assert.NoError(t, err)
	
	cart := &model.Cart{UserId: userID}
	cart, err = s.Cart().CreateCart(userID, cart)
	assert.NoError(t, err)
	
	var comboId int
	err = db.QueryRow(
		"INSERT INTO combos (title, price) VALUES ($1, $2) RETURNING id", 
		"Test Combo", 
		500,
	).Scan(&comboId)
	assert.NoError(t, err)
	
	err = s.Cart().AddCombo(comboId, cart.Id)
	assert.NoError(t, err)
	
	err = s.Cart().DeleteCombo(comboId, cart.Id)
	assert.NoError(t, err)

	var amount int
	err = db.QueryRow(
		"SELECT amount FROM combos_in_cart WHERE cart_id = $1 AND combo_id = $2", 
		cart.Id, 
		comboId,
	).Scan(&amount)
	assert.Error(t, err)

	err = s.Cart().AddCombo(comboId, cart.Id)
	assert.NoError(t, err)
	err = s.Cart().AddCombo(comboId, cart.Id)
	assert.NoError(t, err)

	err = s.Cart().DeleteCombo(comboId, cart.Id)
	assert.NoError(t, err)

	err = db.QueryRow(
		"SELECT amount FROM combos_in_cart WHERE cart_id = $1 AND combo_id = $2", 
		cart.Id, 
		comboId,
	).Scan(&amount)
	assert.NoError(t, err)
	assert.Equal(t, 1, amount)
}

func TestCartRepository_Refresh(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("carts", "products_in_cart", "products", "combos_in_cart", "combos", "users")

	s := sqlstorage.NewSQLStorage(db)
	
	userID, err := createTestUser(db)
	assert.NoError(t, err)
	
	cart := &model.Cart{UserId: userID}
	cart, err = s.Cart().CreateCart(userID, cart)
	assert.NoError(t, err)
	
	var productId int
	err = db.QueryRow(
		"INSERT INTO products (title, price) VALUES ($1, $2) RETURNING id", 
		"Test Product", 
		100,
	).Scan(&productId)
	assert.NoError(t, err)
	
	var comboId int
	err = db.QueryRow(
		"INSERT INTO combos (title, price) VALUES ($1, $2) RETURNING id", 
		"Test Combo", 
		500,
	).Scan(&comboId)
	assert.NoError(t, err)
	
	err = s.Cart().AddProduct(productId, "1 шт", cart.Id)
	assert.NoError(t, err)
	
	err = s.Cart().AddCombo(comboId, cart.Id)
	assert.NoError(t, err)
	
	err = s.Cart().Refresh(cart.Id)
	assert.NoError(t, err)
	
	var productCount int
	err = db.QueryRow(
		"SELECT COUNT(*) FROM products_in_cart WHERE cart_id = $1", 
		cart.Id,
	).Scan(&productCount)
	assert.NoError(t, err)
	assert.Equal(t, 0, productCount)
	
	var comboCount int
	err = db.QueryRow(
		"SELECT COUNT(*) FROM combos_in_cart WHERE cart_id = $1", 
		cart.Id,
	).Scan(&comboCount)
	assert.NoError(t, err)
	assert.Equal(t, 0, comboCount)
}


// func setupTestDB(t *testing.T) (*sql.DB, func(tables ...string)) {
// 	db, err := sql.Open("postgres", databaseURL)
// 	assert.NoError(t, err)
	
// 	err = db.Ping()
// 	assert.NoError(t, err)

// 	teardown := func(tables ...string) {
// 		for _, table := range tables {
// 			_, err := db.Exec("DELETE FROM " + table)
// 			if err != nil {
// 				t.Logf("Ошибка очистки таблицы %s: %v", table, err)
// 			}
// 		}
// 		db.Close()
// 	}
	
// 	return db, teardown
// }

func createTestUser(db *sql.DB) (int, error) {
	var userID int
	err := db.QueryRow(
		"INSERT INTO users (phone) VALUES ($1) RETURNING id",
		"+79155182245",
	).Scan(&userID)
	
	return userID, err
}

// func init() {
// 	db, err := sql.Open("postgres", databaseURL)
// 	if err == nil {
// 		defer db.Close()
		
// 		createTables := []string{
// 			`CREATE TABLE IF NOT EXISTS users (
// 				id SERIAL PRIMARY KEY,
// 				email VARCHAR(255) NOT NULL UNIQUE,
// 				password_hash VARCHAR(255) NOT NULL,
// 				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// 			)`,
// 			`CREATE TABLE IF NOT EXISTS carts (
// 				id SERIAL PRIMARY KEY,
// 				user_id INTEGER NOT NULL REFERENCES users(id),
// 				total_count INTEGER DEFAULT 0
// 			)`,
// 			`CREATE TABLE IF NOT EXISTS products (
// 				id SERIAL PRIMARY KEY,
// 				title VARCHAR(255) NOT NULL,
// 				price INTEGER NOT NULL
// 			)`,
// 			`CREATE TABLE IF NOT EXISTS products_in_cart (
// 				cart_id INTEGER NOT NULL REFERENCES carts(id),
// 				product_id INTEGER NOT NULL REFERENCES products(id),
// 				amount VARCHAR(50) NOT NULL,
// 				count INTEGER DEFAULT 1,
// 				PRIMARY KEY (cart_id, product_id, amount)
// 			)`,
// 			`CREATE TABLE IF NOT EXISTS combos (
// 				id SERIAL PRIMARY KEY,
// 				title VARCHAR(255) NOT NULL,
// 				price INTEGER NOT NULL
// 			)`,
// 			`CREATE TABLE IF NOT EXISTS combos_in_cart (
// 				cart_id INTEGER NOT NULL REFERENCES carts(id),
// 				combo_id INTEGER NOT NULL REFERENCES combos(id),
// 				amount INTEGER DEFAULT 1,
// 				PRIMARY KEY (cart_id, combo_id)
// 			)`,
// 		}
		
// 		for _, query := range createTables {
// 			db.Exec(query)
// 		}
// 	}
// }