package sql_storage

import (
	"github.com/go-next-pizza/internal/app/model"
)

type CartRepository struct {
	storage *SQLStorage
}

func (cr *CartRepository) CreateCart(userId int, c *model.Cart) (*model.Cart, error) {
	c.UserId = userId

	if err := cr.storage.db.QueryRow("INSERT INTO carts (user_id) VALUES ($1) RETURNING id", c.UserId).Scan(&c.Id); err != nil {
		return nil, err
	}

	return c, nil
}

func (cr *CartRepository) AddProduct(productId int, c *model.Cart) error {
	var productCount int

	if err := cr.storage.db.QueryRow("SELECT COUNT(*) FROM products_in_cart WHERE cart_id = $1 AND product_id = $2", c.Id, productId).Scan(&productCount); err != nil {
		return err
	}

	if productCount == 0 {
		_, err := cr.storage.db.Exec("INSERT INTO products_in_cart (cart_id, product_id, amount) VALUES ($1, $2, 1)", c.Id, productId)
		return err
	}

	_, err := cr.storage.db.Exec("UPDATE products_in_cart SET amount = amount + 1 WHERE cart_id = $1 AND product_id = $2", c.Id, productId)

	return err

}

func (cr *CartRepository) AddCombo(comboId int, c *model.Cart) error {
	var comboCount int

	if err := cr.storage.db.QueryRow("SELECT COUNT(*) FROM combos_in_cart WHERE cart_id = $1 AND combo_id = $2", c.Id, comboId).Scan(&comboCount); err != nil {
		return err
	}

	if comboCount == 0 {
		_, err := cr.storage.db.Exec("INSERT INTO combos_in_cart (cart_id, combo_id, amount) VALUES ($1, $2, 1)", c.Id, comboId)
		return err
	}

	_, err := cr.storage.db.Exec("UPDATE combos_in_cart SET amount = amount + 1 WHERE cart_id = $1 AND combo_id = $2", c.Id, comboId)

	return err
}

func (cr *CartRepository) DeleteProduct(productId int, c *model.Cart) error {
	var productCount int

	if err := cr.storage.db.QueryRow("SELECT amount FROM products_in_cart WHERE cart_id = $1 AND product_id = $2", c.Id, productId).Scan(&productCount); err != nil {
		return err
	}

	if productCount == 1 {
			_, err := cr.storage.db.Exec("DELETE FROM products_in_cart WHERE cart_id = $1 AND product_id = $2", c.Id, productId)
			return err
		}

		_, err := cr.storage.db.Exec("UPDATE products_in_cart SET amount = amount - 1 WHERE cart_id = $1 AND product_id = $2", c.Id, productId)

	return err

}

func (cr *CartRepository) DeleteCombo(comboId int, c *model.Cart) error {
	var comboCount int

	if err := cr.storage.db.QueryRow("SELECT amount FROM combos_in_cart WHERE cart_id = $1 AND combo_id = $2", c.Id, comboId).Scan(&comboCount); err != nil {
		return err
	}

	if comboCount == 1 {
			_, err := cr.storage.db.Exec("DELETE FROM combos_in_cart WHERE cart_id = $1 AND combo_id = $2", c.Id, comboId)
			return err
		}

		_, err := cr.storage.db.Exec("UPDATE combos_in_cart SET amount = amount - 1 WHERE cart_id = $1 AND combo_id = $2", c.Id, comboId)

	return err
}

func (cr *CartRepository) Refresh(c *model.Cart) error {
	if _, err := cr.storage.db.Exec("DELETE FROM products_in_cart WHERE cart_id = $1", c.Id); err != nil {
		return err
	}

	if _, err := cr.storage.db.Exec("DELETE FROM combos_in_cart WHERE cart_id = $1", c.Id); err != nil {
		return err
	}

	return nil
}