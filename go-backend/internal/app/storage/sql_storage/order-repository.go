package sql_storage

import (
	"github.com/go-next-pizza/internal/app/model"
)

type OrderRepository struct {
	storage *SQLStorage
}

func (or *OrderRepository) CreateOrder(c *model.Cart, o *model.Order) (*model.Order, error) {

	if err := or.storage.db.QueryRow("INSERT INTO orders (user_id, payment_method, delivery_address, delivery_time, status, total_price, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", o.UserId, o.PaymentMethod, o.DeliveryAddress, o.DeliveryTime, o.Status, o.TotalPrice, o.CreatedAt).Scan(&o.Id); err != nil {
		return nil, err
	}

	rows, err := or.storage.db.Query("SELECT product_id, amount FROM products_in_cart WHERE cart_id = $1", c.Id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	
	for rows.Next() {
		var productId int
		var amount int
		if err := rows.Scan(&productId, &amount); err != nil {
			return nil, err
		}
		if err := or.AddProduct(productId, amount, o); err != nil {
			return nil, err
		}
	}
	
	rows, err = or.storage.db.Query("SELECT combo_id, amount FROM combos_in_cart WHERE cart_id = $1", c.Id)
	
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	
	
	for rows.Next() {
		var comboId int
		var amount int

		if err := rows.Scan(&comboId, &amount); err != nil {
			return nil, err
		}
		
		if err := or.AddCombo(comboId, amount, o); err != nil {
			return nil, err
		}
		
	}

	

	return o, nil
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

func (or *OrderRepository) GetOrderById(id int) (*model.Order, error) {

	return nil, nil
}

func (or *OrderRepository) GetOrdersByUserId(userId int) ([]*model.Order, error) {
	return nil, nil
}
