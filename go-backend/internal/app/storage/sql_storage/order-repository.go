package sql_storage

import (
	"github.com/go-next-pizza/internal/app/model"
)

type OrderRepository struct {
	storage *SQLStorage
}

// func (or *OrderRepository) CreateOrder(c *model.Cart, o *model.Order) (*model.Order, error) {
// 	tx, err := or.storage.db.Begin()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer tx.Rollback()

// 	if err := tx.QueryRow("INSERT INTO orders (user_id, payment_method, delivery_address, delivery_time, status, total_price, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", 
// 		o.UserId, o.PaymentMethod, o.DeliveryAddress, o.DeliveryTime, o.Status, o.TotalPrice, o.CreatedAt).Scan(&o.Id); err != nil {
// 		return nil, err
// 	}

// 	var totalPrice int

// 	rows, err := tx.Query("SELECT product_id, amount FROM products_in_cart WHERE cart_id = $1", c.Id)
// 	if err != nil {
// 		return nil, err
// 	}
	
// 	defer rows.Close()
	
// 	for rows.Next() {
// 		var productId int
// 		var amount int
// 		if err := rows.Scan(&productId, &amount); err != nil {
// 			return nil, err
// 		}

// 		product, err := or.storage.Product().GetProductById(productId)
// 		if err != nil {
// 			return nil, err
// 		}

// 		if _, err := tx.Exec("INSERT INTO products_in_order (order_id, product_id, amount) VALUES ($1, $2, $3)", 
// 			o.Id, productId, amount); err != nil {
// 			return nil, err
// 		}

// 		if product.Price.Valid {
// 			totalPrice += int(product.Price.Int64) * amount
// 		}
// 	}
	
// 	rows, err = tx.Query("SELECT combo_id, amount FROM combos_in_cart WHERE cart_id = $1", c.Id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
	
// 	for rows.Next() {
// 		var comboId int
// 		var amount int

// 		if err := rows.Scan(&comboId, &amount); err != nil {
// 			return nil, err
// 		}

// 		combo, err := or.storage.Combo().GetComboById(comboId)
// 		if err != nil {
// 			return nil, err
// 		}
		
// 		if _, err := tx.Exec("INSERT INTO combos_in_order (order_id, combo_id, amount) VALUES ($1, $2, $3)", 
// 			o.Id, comboId, amount); err != nil {
// 			return nil, err
// 		}

// 		totalPrice += combo.Price
// 	}
	
// 	if _, err := tx.Exec("UPDATE orders SET total_price = $1 WHERE id = $2", totalPrice, o.Id); err != nil {
// 		return nil, err
// 	}

// 	if _, err := tx.Exec("DELETE FROM products_in_cart WHERE cart_id = $1", c.Id); err != nil {
// 		return nil, err
// 	}
// 	if _, err := tx.Exec("DELETE FROM combos_in_cart WHERE cart_id = $1", c.Id); err != nil {
// 		return nil, err
// 	}

// 	if err := tx.Commit(); err != nil {
// 		return nil, err
// 	}

// 	o.TotalPrice = totalPrice
// 	return o, nil
// }

// id bigserial not null primary key,
//   user_id bigint not null,
//   payment_method varchar,
//   delivery_address varchar,
//   delivery_time timestamp,
//   status varchar,
//   total_price int not null,
//   created_at timestamp not null

func (or *OrderRepository) CreateOrder(o *model.Order) (error) {
	_, err := or.storage.db.Exec("INSERT INTO orders (user_id, total_price, created_at) VALUES ($1, $2, $3)", o.UserId, o.TotalPrice, o.CreatedAt)

	return err
}

func (or *OrderRepository) AddProduct(productId int, amount int, o *model.Order) error {

	if _, err := or.storage.db.Exec("INSERT INTO products_in_order (order_id, product_id, amount) VALUES ($1, $2, $3)", o.Id, productId, amount); err != nil {
		return err
	}

	return nil
}

func (or *OrderRepository) AddCombo(comboId int, amount int, o *model.Order) error {

	if _, err := or.storage.db.Exec("INSERT INTO combos_in_order (order_id, combo_id, amount) VALUES ($1, $2, $3)", o.Id, comboId, amount); err != nil {
		return err
	}

	return nil
}

func (or *OrderRepository) GetOrderById(orderId int) (*model.Order, error) {
	order := &model.Order{}

	if err := or.storage.db.QueryRow("SELECT * FROM orders WHERE id = $1", orderId).Scan(
							&order.Id,
							&order.UserId,
							&order.Status, 
							&order.TotalPrice, 
							&order.PaymentMethod, 
							&order.DeliveryAddress,
							&order.CreatedAt,
							&order.DeliveryTime,
						); err != nil {
							return nil, err
						}

	return nil, nil
}

func (or *OrderRepository) GetOrderProducts(orderId int) ([]*model.Product, error) {

	rows, err := or.storage.db.Query("SELECT * FROM products_in_order WHERE order_id = $1", orderId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*model.Product

	for rows.Next() {
		var productId int
		var amount int

		if err := rows.Scan(&productId, &amount); err != nil {
			return nil, err
		}

		product, err := or.storage.Product().GetProductById(productId)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (or *OrderRepository) GetOrderCombos(orderId int) ([]*model.Combo, error) {

	rows, err := or.storage.db.Query("SELECT * FROM combos_in_order WHERE order_id = $1", orderId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var combos []*model.Combo

	for rows.Next() {
		var comboId int
		var amount int

		if err := rows.Scan(&comboId, &amount); err != nil {
			return nil, err
		}

		combo, err := or.storage.Combo().GetComboById(comboId)
		if err != nil {
			return nil, err
		}

		combos = append(combos, combo)
	}

	return combos, nil	
}
