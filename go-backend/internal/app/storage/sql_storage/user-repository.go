package sql_storage

import (
	"database/sql"

	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/storage"
)

type UserRepository struct {
	storage *SQLStorage
}

func (ur *UserRepository) CreateUser(u *model.User) (*model.User, error) {
	tx, err := ur.storage.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}

	if err := tx.QueryRow("INSERT INTO users (email, encrypted_password, name, phone) VALUES ($1, $2, $3, $4) RETURNING id", 
		u.Email, u.EncryptedPassword, u.Name, u.Phone).Scan(&u.Id); err != nil {
		return nil, err
	}

	u.Cart = &model.Cart{}
	if err := tx.QueryRow("INSERT INTO carts (user_id) VALUES ($1) RETURNING id", u.Id).Scan(&u.Cart.Id); err != nil {
		return nil, err
	}
	u.Cart.UserId = u.Id

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	u.Cart = &model.Cart{}

	if err := ur.storage.db.QueryRow("SELECT id, email, encrypted_password, name, phone FROM users WHERE email = $1", email).Scan(&u.Id, &u.Email, &u.EncryptedPassword, &u.Name, &u.Phone); err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.ErrRecordNotFound
		}
		
		return nil, err
	}

	return u, nil
}

func (ur *UserRepository) FindById(id int) (*model.User, error) {
	u := &model.User{}
	u.Cart = &model.Cart{}

	if err := ur.storage.db.QueryRow("SELECT id, email, encrypted_password, name, phone FROM users WHERE id = $1", id).Scan(&u.Id, &u.Email, &u.EncryptedPassword, &u.Name, &u.Phone); err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.ErrRecordNotFound
		}
		
		return nil, err
	}

	return u, nil
}

func (ur *UserRepository) GetOrders(userId int) ([]*model.Order, error) {
	rows, err := ur.storage.db.Query("SELECT id, user_id, payment_method, delivery_address, delivery_time, status, total_price, created_at FROM orders WHERE user_id = $1", userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []*model.Order

	for rows.Next() {
		order := &model.Order{}

		if err := rows.Scan(&order.Id, &order.UserId, &order.PaymentMethod, &order.DeliveryAddress, &order.DeliveryTime, &order.Status, &order.TotalPrice, &order.CreatedAt); err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

// func (ur *UserRepository) GetCart(userId int) (*model.Cart, error) {

// 	var cartId int

// 	if err := ur.storage.db.QueryRow("SELECT id FROM carts WHERE user_id = $1", userId).Scan(&cartId); err != nil {
// 		return nil, err
// 	}

// 	cart := &model.Cart{
// 		Id: cartId,
// 	}

// 	rows, err := ur.storage.db.Query("SELECT product_id, COUNT(product_id) FROM products_in_cart WHERE cart_id = $1 GROUP BY product_id", cart.Id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for rows.Next() {
// 		var productId int
// 		var amount int

// 		if err := rows.Scan(&productId, &amount); err != nil {
// 			return nil, err
// 		}

// 		cart.Products[productId] = amount
// 	}

// 	rows, err = ur.storage.db.Query("SELECT combo_id, COUNT(combo_id) FROM combos_in_cart WHERE cart_id = $1 GROUP BY combo_id", cart.Id)
// 	if err != nil {
// 		return nil, err
// 	}
	
// 	for rows.Next() {
// 		var comboId int
// 		var amount int

// 		if err := rows.Scan(&comboId, &amount); err != nil {
// 			return nil, err
// 		}

// 		cart.Combos[comboId] = amount
// 	}

// 	return cart, nil
// }