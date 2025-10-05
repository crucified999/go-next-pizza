package sql_storage

import (
	"database/sql"

	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/storage"
)

type IngredientRepository struct {
	storage *SQLStorage
}

func (ir *IngredientRepository) GetIngredientByID(id int) (*model.Ingredient, error) {
	ingredient := &model.Ingredient{}
	
	if err := ir.storage.db.QueryRow(
		"SELECT id, title, product_id, replacable FROM pizza_ingredients WHERE id = $1",
		id,
	).Scan(&ingredient.ID, &ingredient.Title, &ingredient.ProductID, &ingredient.Replacable); err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.ErrRecordNotFound
		}
		return nil, err
	}

	// Получаем цену ингредиента из связанного продукта
	if err := ir.storage.db.QueryRow(
		"SELECT price FROM products WHERE id = $1",
		ingredient.ProductID,
	).Scan(&ingredient.Price); err != nil {
		return nil, err
	}

	return ingredient, nil
}

func (ir *IngredientRepository) GetIngredients() ([]*model.Ingredient, error) {
	rows, err := ir.storage.db.Query(
		"SELECT id, title, product_id, replacable FROM pizza_ingredients ORDER BY title",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ingredients []*model.Ingredient
	for rows.Next() {
		ingredient := &model.Ingredient{}
		if err := rows.Scan(&ingredient.ID, &ingredient.Title, &ingredient.ProductID, &ingredient.Replacable); err != nil {
			return nil, err
		}
		
		// Получаем цену ингредиента из связанного продукта
		if err := ir.storage.db.QueryRow(
			"SELECT price FROM products WHERE id = $1",
			ingredient.ProductID,
		).Scan(&ingredient.Price); err != nil {
			return nil, err
		}
		
		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil
}