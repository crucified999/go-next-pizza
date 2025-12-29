package sql_storage

import (
	"database/sql"
	"log"

	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/storage"
)

type ProductRepository struct {
	storage *SQLStorage
}

func (pr *ProductRepository) GetProductById(id int) (*model.Product, error) {
	product := &model.Product{}

	if err := pr.storage.db.QueryRow(
		"SELECT id, title, description, price, image, weight, amount FROM products WHERE id = $1",
		id,
	).Scan(&product.Id, &product.Title, &product.Description, &product.Price, &product.Image, &product.Weight, &product.Amount); err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.ErrRecordNotFound
		}
		return nil, err
	}

	return product, nil
}

func (pr *ProductRepository) GetProducts() ([]*model.Product, error) {
	rows, err := pr.storage.db.Query(
		"SELECT id, title, description, price, image, weight, amount FROM products ORDER BY title",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*model.Product
	for rows.Next() {
		product := &model.Product{}
		if err := rows.Scan(&product.Id, &product.Title, &product.Description, &product.Price, &product.Image, &product.Weight, &product.Amount); err != nil {
			log.Printf("Ошибка получения продуктов")
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (pr *ProductRepository) GetProductsByCategory(cId int) ([]*model.Product, error) {

	rows, err := pr.storage.db.Query(
			`SELECT products.id, title, description, price, image, weight, amount 
			FROM products JOIN categories ON product_id = products.id
			WHERE categories.category_id = $1
		 `, cId,
		)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*model.Product

	for rows.Next() {
		product := &model.Product{}

		if err := rows.Scan(&product.Id, &product.Title, &product.Description, &product.Price, &product.Image, &product.Weight, &product.Amount); err != nil {
			log.Printf("Ошибка получения продуктов по категории")
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (pr *ProductRepository) GetProductCategory(productId int) (string, error) {

	var categoryId int

	if err := pr.storage.db.QueryRow("SELECT category_id FROM categories WHERE product_id = $1", productId).Scan(&categoryId); err != nil {
		return "", nil
	}

	return pr.ConvertIdToCategory(categoryId), nil
}

func (pr *ProductRepository) GetProductsVariants(productId int) ([]*model.ProductVariant, error) {

	rows, err := pr.storage.db.Query("SELECT product_id, size, dough_type, image, weight, price FROM products_with_size WHERE product_id = $1", productId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var variants []*model.ProductVariant

	for rows.Next() {
		variant := &model.ProductVariant{}

		if err := rows.Scan(&variant.ProductId, &variant.Size, &variant.DoughType, &variant.Image, &variant.Weight, &variant.Price); err != nil {
			return nil, err
		}

		variants = append(variants, variant)
	}

	return variants, nil
}

func (pr *ProductRepository) GetProductIngredients(productId int) ([]*model.Ingredient, error) {
	rows, err := pr.storage.db.Query("SELECT id, title, product_id, replacable FROM pizza_ingredients WHERE product_id = $1", productId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var ingredients []*model.Ingredient

	for rows.Next() {
		ingredient := &model.Ingredient{}

		if err := rows.Scan(&ingredient.Id, &ingredient.Title, &ingredient.ProductID, &ingredient.Replacable); err != nil {
			return nil, err
		}

		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil
}

func (pr *ProductRepository) GetProductToppings(productId int) ([]*model.Topping, error) {
	rows, err := pr.storage.db.Query("SELECT id, title, product_id, image, price FROM toppings WHERE product_id = $1", productId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var toppings []*model.Topping

	for rows.Next() {
		topping := &model.Topping{}
		if err := rows.Scan(&topping.Id, &topping.Title, &topping.ProductID, &topping.Image, &topping.Price); err != nil {
			return nil, err
		}
		
    toppings = append(toppings, topping)
	}

	return toppings, nil
}

func (pr *ProductRepository) GetProductVariant(productId int, size string) (*model.ProductVariant, error) {
	pv := &model.ProductVariant{
		ProductId: productId,
	}

	if err := pr.storage.db.QueryRow("SELECT size, image, weight, price FROM products_with_size WHERE product_id = $1 AND size = $2", productId, size).Scan(&pv.Size, &pv.Image, &pv.Weight, &pv.Price); err != nil {
		if err == sql.ErrNoRows {
			err = pr.storage.db.QueryRow("SELECT title, image, amount, weight, price FROM products WHERE id = $1", pv.ProductId).Scan(&pv.Title, &pv.Image, &pv.Size, &pv.Weight, &pv.Price)

			return pv, err
		}
	}

	if err := pr.storage.db.QueryRow("SELECT title FROM products WHERE id = $1", pv.ProductId).Scan(&pv.Title); err != nil {
		return nil, err
	}

	return pv, nil
}

func (pr *ProductRepository) GetPizzaVariant(pizza *model.PizzaVariant) (*model.PizzaVariant, error) {
	pv := &model.PizzaVariant{
		PizzaId: pizza.PizzaId,
	}

	if err := pr.storage.db.QueryRow("SELECT dough_type, size, image, weight, price FROM products_with_size WHERE product_id = $1 AND size = $2 AND dough_type = $3", pizza.PizzaId, pizza.Size, pizza.Dough).Scan(&pv.Dough, &pv.Size, &pv.Image, &pv.Weight, &pv.Price); err != nil {
		return nil, err
	}

	if err := pr.storage.db.QueryRow("SELECT title FROM products WHERE id = $1", pv.PizzaId).Scan(&pv.Title); err != nil {
		return nil, err
	}

	// rows, err := pr.storage.db.Query("SELECT * FROM toppings WHERE product_id = $1", pv.Id)

	// if err != nil {
	// 	return nil, err
	// }

	// defer rows.Close()

	// for rows.Next() {
	// 	topping := &model.Topping{}

	// 	if err := rows.Scan(&topping.Id, &topping.Title, &topping.Price, &topping.ProductID, &topping.Image); err != nil {
	// 		return nil, err
	// 	}

	// 	pv.Toppings = append(pv.Toppings, topping)
		
	// }

	return pv, nil
}


func (pr *ProductRepository) ConvertIdToCategory(categoryId int) string {

	var category string

	switch categoryId {
	case 1:
		category = "pizza"
	case 2:
		category = "combo"
	case 3:
		category = "snack"
	case 4:
		category = "shake"
	case 5:
		category = "coffee"
	case 6:
		category = "drink"
	case 7:
		category = "dessert"
	case 8:
		category = "sauce"
	}

	return category
}

func (pr *ProductRepository) ConvertCategorToId(category string) int {

	var categoryId int

	switch category {
	case "pizzas":
		categoryId = 1
	case "combos":
		categoryId = 2
	case "snacks":
		categoryId = 3
	case "shakes":
		categoryId = 4
	case "coffee":
		categoryId = 5
	case "drinks":
		categoryId = 6
	case "desserts":
		categoryId = 7
	case "sauces":
		categoryId = 8
	}

	return categoryId
}


func (pr *ProductRepository) isPizza(productId int) (bool) {

	_, err := pr.storage.db.Query(`
		SELECT id FROM pizza_sizes WHERE pizza_id = $1
	`, productId)

	if err != sql.ErrNoRows {
		return true
	}

	return false

}