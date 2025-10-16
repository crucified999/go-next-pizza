package sql_storage

import (
	"database/sql"

	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/storage"
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

func (cr *CartRepository) GetCartByUserId(userId int) (*model.Cart, error) {
	cart := &model.Cart{}

	if err := cr.storage.db.QueryRow("SELECT id, user_id FROM carts WHERE user_id = $1", userId).Scan(&cart.Id, &cart.UserId); err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.ErrRecordNotFound
		}
		return nil, err
	}

	return cart, nil
}

func (cr *CartRepository) GetCartProducts(cartId int) ([]*model.CartProduct, error) {
	rows, err := cr.storage.db.Query("SELECT product_id, amount FROM products_in_cart WHERE cart_id = $1", cartId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var cartProducts []*model.CartProduct

	for rows.Next() {
		var productId int
		var amount int

		if err := rows.Scan(&productId, &amount); err != nil {
			return nil, err
		}

		product, err := cr.storage.Product().GetProductById(productId)
		if err != nil {
			return nil, err
		}

		cartProducts = append(cartProducts, &model.CartProduct{
			Product: product,
			Amount:  amount,
		})
	}

	return cartProducts, nil
}

func (cr *CartRepository) GetCartCombos(cartId int) ([]*model.CartCombo, error) {
	rows, err := cr.storage.db.Query("SELECT combo_id, amount FROM combos_in_cart WHERE cart_id = $1", cartId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var cartCombos []*model.CartCombo

	for rows.Next() {
		var comboId int
		var amount int

		if err := rows.Scan(&comboId, &amount); err != nil {
			return nil, err
		}

		combo, err := cr.storage.Combo().GetComboById(comboId)
		if err != nil {
			return nil, err
		}

		cartCombos = append(cartCombos, &model.CartCombo{
			Combo:  combo,
			Amount: amount,
		})
	}

	return cartCombos, nil
}

func (cr *CartRepository) AddProduct(productId int, cartId int) error {
	var productCount int

	if err := cr.storage.db.QueryRow("SELECT COUNT(*) FROM products_in_cart WHERE cart_id = $1 AND product_id = $2", cartId, productId).Scan(&productCount); err != nil {
		return err
	}

	if productCount == 0 {
		_, err := cr.storage.db.Exec("INSERT INTO products_in_cart (cart_id, product_id, amount) VALUES ($1, $2, 1)", cartId, productId)
		return err
	}

	_, err := cr.storage.db.Exec("UPDATE products_in_cart SET amount = amount + 1 WHERE cart_id = $1 AND product_id = $2", cartId, productId)

	return err
}

func (cr *CartRepository) AddCombo(comboId int, cartId int) error {
	var comboCount int

	if err := cr.storage.db.QueryRow("SELECT COUNT(*) FROM combos_in_cart WHERE cart_id = $1 AND combo_id = $2", cartId, comboId).Scan(&comboCount); err != nil {
		return err
	}

	if comboCount == 0 {
		_, err := cr.storage.db.Exec("INSERT INTO combos_in_cart (cart_id, combo_id, amount) VALUES ($1, $2, 1)", cartId, comboId)
		return err
	}

	_, err := cr.storage.db.Exec("UPDATE combos_in_cart SET amount = amount + 1 WHERE cart_id = $1 AND combo_id = $2", cartId, comboId)

	return err
}

func (cr *CartRepository) DeleteProduct(productId int, cartId int) error {
	var productCount int

	if err := cr.storage.db.QueryRow("SELECT amount FROM products_in_cart WHERE cart_id = $1 AND product_id = $2", cartId, productId).Scan(&productCount); err != nil {
		return err
	}

	if productCount == 1 {
			_, err := cr.storage.db.Exec("DELETE FROM products_in_cart WHERE cart_id = $1 AND product_id = $2", cartId, productId)
			return err
		}

		_, err := cr.storage.db.Exec("UPDATE products_in_cart SET amount = amount - 1 WHERE cart_id = $1 AND product_id = $2", cartId, productId)

	return err

}

func (cr *CartRepository) DeleteCombo(comboId int, cartId int) error {
	var comboCount int

	if err := cr.storage.db.QueryRow("SELECT amount FROM combos_in_cart WHERE cart_id = $1 AND combo_id = $2", cartId, comboId).Scan(&comboCount); err != nil {
		return err
	}

	if comboCount == 1 {
			_, err := cr.storage.db.Exec("DELETE FROM combos_in_cart WHERE cart_id = $1 AND combo_id = $2", cartId, comboId)
			return err
		}

		_, err := cr.storage.db.Exec("UPDATE combos_in_cart SET amount = amount - 1 WHERE cart_id = $1 AND combo_id = $2", cartId, comboId)

	return err
}

func (cr *CartRepository) Refresh(cartId int) error {
	if _, err := cr.storage.db.Exec("DELETE FROM products_in_cart WHERE cart_id = $1", cartId); err != nil {
		return err
	}

	if _, err := cr.storage.db.Exec("DELETE FROM combos_in_cart WHERE cart_id = $1", cartId); err != nil {
		return err
	}

	return nil
}