package model

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	Weight      int     `json:"weight"`
	Amount      int     `json:"amount"`
}

type Ingredient struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	ProductID   int     `json:"product_id"`
	Replacable  int     `json:"replacable"`
	Price       float64 `json:"price"`
}