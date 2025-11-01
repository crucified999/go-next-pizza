package model

import "database/sql"

type Product struct {
	Id          int               `json:"id"`
	Category    string            `json:"category"`
	Title       sql.NullString    `json:"title"`
	Description sql.NullString    `json:"description,omitempty"`
	Price       sql.NullInt64     `json:"price"`
	Image       sql.NullString    `json:"image"`
	Amount      sql.NullFloat64   `json:"amount,omitempty"`
	Weight      sql.NullInt64     `json:"weight"`
	Variants    []*ProductVariant `json:"variants"`
	Ingredients []*Ingredient     `json:"ingredients,omitempty"`
	Toppings    []*Topping        `json:"toppings,omitempty"`
}

type ProductVariant struct {
	ProductId int `json:"productId"`
	DoughType sql.NullInt64 `json:"doughType,omitempty"`
	Size      string `json:"size"`
	Image     string `json:"image"`
	Weight    int `json:"weight"`
	Price     int `json:"price"`
}

type Topping struct {
	Id int `json:"id"`
	Title string `json:"title"`
	ProductID int `json:"product_id"`
	Image string `json:"image"`
	Price int `json:"price"`
}

type Ingredient struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	ProductID   int     `json:"product_id"`
	Replacable  int     `json:"replacable"`
	Price       float64 `json:"price"`
}