package sql_storage

import (
	"database/sql"

	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/storage"
)

type CustomPizzaRepository struct {
	storage *SQLStorage
}

func (cpr *CustomPizzaRepository) CreateCustomPizza(cp *model.CustomPizza) (*model.CustomPizza, error) {
	if err := cpr.storage.db.QueryRow(
		"INSERT INTO custom_pizzas (user_id, base_pizza_id, name, total_price) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at",
		cp.UserID, cp.BasePizzaID, cp.Name, cp.TotalPrice,
	).Scan(&cp.ID, &cp.CreatedAt, &cp.UpdatedAt); err != nil {
		return nil, err
	}

	for _, ingredient := range cp.Ingredients {
		ingredient.CustomPizzaID = cp.ID
		if err := cpr.addIngredient(&ingredient); err != nil {
			return nil, err
		}
	}

	return cp, nil
}

func (cpr *CustomPizzaRepository) GetCustomPizzaByID(id int) (*model.CustomPizza, error) {
	cp := &model.CustomPizza{}
	
	if err := cpr.storage.db.QueryRow(
		"SELECT id, user_id, base_pizza_id, name, total_price, created_at, updated_at FROM custom_pizzas WHERE id = $1",
		id,
	).Scan(&cp.ID, &cp.UserID, &cp.BasePizzaID, &cp.Name, &cp.TotalPrice, &cp.CreatedAt, &cp.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.ErrRecordNotFound
		}
		return nil, err
	}

	// Загружаем ингредиенты
	ingredients, err := cpr.getIngredientsByCustomPizzaID(cp.ID)
	if err != nil {
		return nil, err
	}
	cp.Ingredients = ingredients

	return cp, nil
}

func (cpr *CustomPizzaRepository) GetCustomPizzasByUserID(userID int) ([]*model.CustomPizza, error) {
	rows, err := cpr.storage.db.Query(
		"SELECT id, user_id, base_pizza_id, name, total_price, created_at, updated_at FROM custom_pizzas WHERE user_id = $1 ORDER BY created_at DESC",
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customPizzas []*model.CustomPizza
	for rows.Next() {
		cp := &model.CustomPizza{}
		if err := rows.Scan(&cp.ID, &cp.UserID, &cp.BasePizzaID, &cp.Name, &cp.TotalPrice, &cp.CreatedAt, &cp.UpdatedAt); err != nil {
			return nil, err
		}

		// Загружаем ингредиенты
		ingredients, err := cpr.getIngredientsByCustomPizzaID(cp.ID)
		if err != nil {
			return nil, err
		}
		cp.Ingredients = ingredients

		customPizzas = append(customPizzas, cp)
	}

	return customPizzas, nil
}

func (cpr *CustomPizzaRepository) UpdateCustomPizza(cp *model.CustomPizza) (*model.CustomPizza, error) {
	// Обновляем основную информацию
	if _, err := cpr.storage.db.Exec(
		"UPDATE custom_pizzas SET name = $1, total_price = $2, updated_at = NOW() WHERE id = $3",
		cp.Name, cp.TotalPrice, cp.ID,
	); err != nil {
		return nil, err
	}

	// Удаляем старые ингредиенты
	if _, err := cpr.storage.db.Exec(
		"DELETE FROM custom_pizza_ingredients WHERE custom_pizza_id = $1",
		cp.ID,
	); err != nil {
		return nil, err
	}

	// Добавляем новые ингредиенты
	for _, ingredient := range cp.Ingredients {
		ingredient.CustomPizzaID = cp.ID
		if err := cpr.addIngredient(&ingredient); err != nil {
			return nil, err
		}
	}

	return cp, nil
}

func (cpr *CustomPizzaRepository) DeleteCustomPizza(id int) error {
	_, err := cpr.storage.db.Exec("DELETE FROM custom_pizzas WHERE id = $1", id)
	return err
}

func (cpr *CustomPizzaRepository) addIngredient(ingredient *model.CustomPizzaIngredient) error {
	return cpr.storage.db.QueryRow(
		"INSERT INTO custom_pizza_ingredients (custom_pizza_id, ingredient_id, is_added) VALUES ($1, $2, $3) RETURNING id, created_at",
		ingredient.CustomPizzaID, ingredient.IngredientID, ingredient.IsAdded,
	).Scan(&ingredient.ID, &ingredient.CreatedAt)
}

func (cpr *CustomPizzaRepository) getIngredientsByCustomPizzaID(customPizzaID int) ([]model.CustomPizzaIngredient, error) {
	rows, err := cpr.storage.db.Query(
		"SELECT id, custom_pizza_id, ingredient_id, is_added, created_at FROM custom_pizza_ingredients WHERE custom_pizza_id = $1",
		customPizzaID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ingredients []model.CustomPizzaIngredient
	for rows.Next() {
		ingredient := model.CustomPizzaIngredient{}
		if err := rows.Scan(&ingredient.ID, &ingredient.CustomPizzaID, &ingredient.IngredientID, &ingredient.IsAdded, &ingredient.CreatedAt); err != nil {
			return nil, err
		}
		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil
}