package model

import (
	"time"
)

type CustomPizza struct {
	ID          int                      `json:"id"`
	UserID      int                      `json:"user_id"`
	BasePizzaID int                      `json:"base_pizza_id"`
	Name        string                   `json:"name"`
	TotalPrice  float64                  `json:"total_price"`
	Ingredients []CustomPizzaIngredient  `json:"ingredients"`
	CreatedAt   time.Time                `json:"created_at"`
	UpdatedAt   time.Time                `json:"updated_at"`
}

type CustomPizzaIngredient struct {
	ID           int     `json:"id"`
	CustomPizzaID int     `json:"custom_pizza_id"`
	IngredientID int     `json:"ingredient_id"`
	IsAdded      bool    `json:"is_added"`
	CreatedAt    time.Time `json:"created_at"`
}

type CreateCustomPizzaRequest struct {
	BasePizzaID int                           `json:"base_pizza_id"`
	Name        string                        `json:"name"`
	Ingredients []CustomPizzaIngredientRequest `json:"ingredients"`
}

type CustomPizzaIngredientRequest struct {
	IngredientID int    `json:"ingredient_id"`
	Action       string `json:"action"`
}

type UpdateCustomPizzaRequest struct {
	Name        string                        `json:"name"`
	Ingredients []CustomPizzaIngredientRequest `json:"ingredients"`
}