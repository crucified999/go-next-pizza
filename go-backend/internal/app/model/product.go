package model

import "database/sql"

type Product struct {
	Id          int             `json:"id"`
	Category    string          `json:"category"`
	Title       sql.NullString  `json:"title"`
	Description sql.NullString  `json:"description,omitempty"`
	Price       sql.NullInt64   `json:"price"`
	Image       sql.NullString  `json:"image"`
	Amount      sql.NullFloat64 `json:"amount,omitempty"`
	Weight      sql.NullInt64   `json:"weight"`
}

type Ingredient struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	ProductID   int     `json:"product_id"`
	Replacable  int     `json:"replacable"`
	Price       float64 `json:"price"`
}