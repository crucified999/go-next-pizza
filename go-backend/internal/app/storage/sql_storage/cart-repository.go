package sql_storage

import (
	"database/sql"
	"log"

	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/storage"
)

type CartRepository struct {
	storage *SQLStorage
}

func (cr *CartRepository) CreateCart(userId int, c *model.Cart) (*model.Cart, error) {
	c.UserId = userId

	if err := cr.storage.db.QueryRow("INSERT INTO carts (user_id, total_count) VALUES ($1, 0) RETURNING id", c.UserId).Scan(&c.Id); err != nil {
		return nil, err
	}

	return c, nil
}

func (cr *CartRepository) GetCartByUserId(userId int) (*model.Cart, error) {
	cart := &model.Cart{}

	if err := cr.storage.db.QueryRow("SELECT id, user_id, total_count FROM carts WHERE user_id = $1", userId).Scan(&cart.Id, &cart.UserId, &cart.TotalCount); err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.ErrRecordNotFound
		}
		return nil, err
	}

	return cart, nil
}

func (cr *CartRepository) GetCartProducts(cartId int) ([]*model.CartProduct, error) {
	rows, err := cr.storage.db.Query("SELECT product_id, amount, count FROM products_in_cart WHERE cart_id = $1", cartId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var cartProducts []*model.CartProduct

	for rows.Next() {
		var productId int
		var amount string
		var count int

		if err := rows.Scan(&productId, &amount, &count); err != nil {
			return nil, err
		}

		product, err := cr.storage.Product().GetProductVariant(productId, amount)
		if err != nil {
			return nil, err
		}

		cartProducts = append(cartProducts, &model.CartProduct{
			Product: product,
			Count: count,
		})
	}

	return cartProducts, nil
}

func (cr *CartRepository) GetCartToppings(mask int) ([]*model.Topping, error) {
	tids := cr.storage.cartRepository.maskToToppings(mask)

	var toppings []*model.Topping

	for _, t := range tids {
		topping := &model.Topping{
			Id: int(t),
		}

		if err := cr.storage.db.QueryRow("SELECT title, price FROM toppings WHERE id = $1", t).Scan(&topping.Title, &topping.Price); err != nil {
			return nil, err
		}

		toppings = append(toppings, topping)
	}

	return toppings, nil
}

func (cr *CartRepository) GetCartPizzas(cartId int) ([]*model.CartPizza, error) {
	rows, err := cr.storage.db.Query("SELECT pizza_id, count, size, dough, toppings_mask FROM pizzas_in_cart WHERE cart_id = $1", cartId)

	if err != nil {
		log.Print("Ошибка получения строк", err)
		return nil, err
	}

	defer rows.Close()

	var cartPizzas []*model.CartPizza

	for rows.Next() {
		pizza := &model.PizzaVariant{}
		var mask int
		var count int

		if err := rows.Scan(&pizza.PizzaId, &count, &pizza.Size, &pizza.Dough, &mask); err != nil {
			return nil, err
		}

		pizza, err := cr.storage.Product().GetPizzaVariant(pizza)

		if err != nil {
			return nil, err
		}

		pizza.Toppings, err = cr.storage.cartRepository.GetCartToppings(mask)

		// pizza, err := cr.storage.Product().GetPizzaVariant(&model.PizzaVariant{
		// 	Id: pizzaId,
		// })

		if err != nil {
			return nil, err
		}

		cartPizzas = append(cartPizzas, &model.CartPizza{
			Pizza: pizza,
			Count: count,
		})
	}

	return cartPizzas, err
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

func (cr *CartRepository) AddProduct(productId int, amount string, cartId int) error {
	var productCount int
	// var totalCount int

	if err := cr.storage.db.QueryRow("SELECT COUNT(*) FROM products_in_cart WHERE cart_id = $1 AND product_id = $2 AND amount = $3", cartId, productId, amount).Scan(&productCount); err != nil {
		return err
	}

	// if err := cr.storage.db.QueryRow("SELECT total_count FROM carts WHERE id = $1", cartId).Scan(&totalCount); err != nil {
	// 	return err
	// }

	if productCount == 0 {
		_, err := cr.storage.db.Exec("INSERT INTO products_in_cart (cart_id, product_id, amount, count) VALUES ($1, $2, $3, 1)", cartId, productId, amount)
		return err
	}

	if _, err := cr.storage.db.Exec("UPDATE products_in_cart SET count = count + 1 WHERE cart_id = $1 AND product_id = $2", cartId, productId); err != nil {
		return err
	}

	// if totalCount == 0 {
	// 	if _, err := cr.storage.db.Exec("UPDATE carts SET total_count = 1 WHERE id = $1", cartId); err != nil {
	// 		return err
	// 	}
	// } else {
	// 	if _, err := cr.storage.db.Exec("UPDATE carts SET total_count = total_counWHEREt + 1  id = $1", cartId); err != nil {
	// 		return err
	// 	}
	// }
	
	return nil
}

func (cr *CartRepository) AddPizza(pizza *model.PizzaVariant, mask int, cartId int) error {
	var productCount int

	if err := cr.storage.db.QueryRow("SELECT COUNT(*) FROM pizzas_in_cart WHERE cart_id = $1 AND pizza_id = $2 AND size = $3 AND dough = $4 AND toppings_mask = $5", cartId, pizza.PizzaId, pizza.Size, pizza.Dough, mask).Scan(&productCount); err != nil {
		return err
	} 

	if productCount == 0 {
		if _, err := cr.storage.db.Exec("INSERT INTO pizzas_in_cart (cart_id, pizza_id, dough, size, toppings_mask, count) VALUES ($1, $2, $3, $4, $5, 1)", cartId,  pizza.PizzaId, pizza.Dough, pizza.Size, mask); err != nil {
			return err
		}
		
		// toppings := cr.maskToToppings(mask)

		
		// for _, i := range toppings {
		// 	if	_, err := cr.storage.db.Exec("INSERT toppings_in_cart (cart_id, product_id, topping_id) VALUES ($1, $2, $3)", cartId, pizza.Id, i); err != nil {
		// 		return err
		// 	}
		// }
		return nil
	}

	_, err := cr.storage.db.Exec("UPDATE pizzas_in_cart SET count = count + 1 WHERE cart_id = $1 AND pizza_id = $2 AND size = $3 AND dough = $4 AND toppings_mask = $5", cartId, pizza.PizzaId, pizza.Size, pizza.Dough, mask)
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

func (cr *CartRepository) DeleteProduct(productId int, cartId int, amount string) error {
	var productCount int

	if err := cr.storage.db.QueryRow("SELECT count FROM products_in_cart WHERE cart_id = $1 AND product_id = $2 AND amount = $3", cartId, productId, amount).Scan(&productCount); err != nil {
		return err
	}

	if productCount == 1 {
			_, err := cr.storage.db.Exec("DELETE FROM products_in_cart WHERE cart_id = $1 AND product_id = $2 AND amount = $3", cartId, productId, amount)
			return err
		}

	if _, err := cr.storage.db.Exec("UPDATE products_in_cart SET count = count - 1 WHERE cart_id = $1 AND product_id = $2 AND amount = $3", cartId, productId, amount); err != nil {
		return err
	}

	if _, err := cr.storage.db.Exec("UPDATE carts SET total_count = total_count - 1 WHERE id = $1", cartId); err != nil {
		return err
	}

	return nil
}

func (cr *CartRepository) DeletePizza(pizza *model.PizzaVariant, mask int, cartId int) error {
	var productCount int

	if err := cr.storage.db.QueryRow("SELECT count FROM pizzas_in_cart WHERE cart_id = $1 AND pizza_id = $2 AND size = $3 AND dough = $4 AND toppings_mask = $5", cartId, pizza.PizzaId, pizza.Size, pizza.Dough, mask).Scan(&productCount); err != nil {
		return err
	}

	log.Print("Пицца нашлась")

	if productCount == 1 {
			_, err := cr.storage.db.Exec("DELETE FROM pizzas_in_cart WHERE cart_id = $1 AND pizza_id = $2 AND size = $3 AND dough = $4 AND toppings_mask = $5", cartId, pizza.PizzaId, pizza.Size, pizza.Dough, mask)
			return err
		}

	if _, err := cr.storage.db.Exec("UPDATE pizzas_in_cart SET count = count - 1 WHERE cart_id = $1 AND pizza_id = $2 AND size = $3 AND dough = $4 AND toppings_mask = $5", cartId, pizza.PizzaId, pizza.Size, pizza.Dough, mask); err != nil {
		return err
	}

	if _, err := cr.storage.db.Exec("UPDATE carts SET total_count = total_count - 1 WHERE id = $1", cartId); err != nil {
		return err
	}

	return nil
}

func (cr *CartRepository) DeleteProductCompletely(productId int, cartId int, amount string) error {
	var productCount int;

	if err := cr.storage.db.QueryRow("SELECT count FROM products_in_cart WHERE cart_id = $1 AND product_id = $2 AND amount = $3", cartId, productId, amount).Scan(&productCount); err != nil {
		return err
	}

	if _, err := cr.storage.db.Exec("DELETE FROM products_in_cart WHERE cart_id = $1 AND product_id = $2 AND amount = $3", cartId, productId, amount); err != nil {
		return err
	}

	if _, err := cr.storage.db.Exec("UPDATE carts SET total_count = total_count - $1 WHERE id = $2", productCount, cartId); err != nil {
		return err
	}

	return nil
}

func (cr *CartRepository) DeletePizzaCompletely(pizza *model.PizzaVariant, mask int, cartId int) error {
	var productCount int

	if err := cr.storage.db.QueryRow("SELECT count FROM pizzas_in_cart WHERE cart_id = $1 AND pizza_id = $2 AND size = $3 AND dough = $4 AND toppings_mask = $5", cartId, pizza.PizzaId, pizza.Size, pizza.Dough, mask).Scan(&productCount); err != nil {
		return err
	}

	if _, err := cr.storage.db.Exec("DELETE FROM pizzas_in_cart WHERE cart_id = $1 AND pizza_id = $2 AND size = $3 AND dough = $4 AND toppings_mask = $5", cartId, pizza.PizzaId, pizza.Size, pizza.Dough, mask); err != nil {
		return err
	}

	if _, err := cr.storage.db.Exec("UPDATE carts SET total_count = total_count - $1 WHERE id = $2", productCount, cartId); err != nil {
		return err
	}

	return nil
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

	if _, err := cr.storage.db.Exec("DELETE FROM pizzas_in_cart WHERE cart_id = $1", cartId); err != nil {
		return err
	}

	if _, err := cr.storage.db.Exec("DELETE FROM combos_in_cart WHERE cart_id = $1", cartId); err != nil {
		return err
	}

	return nil
}

func (cr *CartRepository) toppingsToMask(toppingIDs []int) int {
	var mask int = 0
	for _, id := range toppingIDs {
			if id > 0 && id <= 64 {
					mask |= 1 << (id - 1)
			}
	}

	return mask
}

func (cr *CartRepository) maskToToppings(mask int) []int {
	var toppings []int
	for i := 0; i < 64; i++ {
			if mask&(1<<i) != 0 {
					toppings = append(toppings, i+1)
			}
	}
	return toppings
}
