package model

import "database/sql"

type Product struct {
	Id          int               `json:"id"`
	Category    string            `json:"category"`
	Title       sql.NullString    `json:"title"`
	Description sql.NullString    `json:"description,omitempty"`
	Price       sql.NullInt64     `json:"price"`
	Image       sql.NullString    `json:"image"`
	Amount      sql.NullString    `json:"amount,omitempty"`
	Weight      sql.NullInt64     `json:"weight"`
	Variants    []*ProductVariant `json:"variants"`
	Ingredients []*Ingredient     `json:"ingredients,omitempty"`
	Toppings    []*Topping        `json:"toppings,omitempty"`
}


type Pizza struct {
	Id int `json:"id"`
	Dough string `json:"dough"`
	Size string `json:"size"`
	Ingredients []*Ingredient
}

type PizzaVariant struct {
	PizzaId int `json:"pizzaId"`
	Title string `json:"title"`
	Dough int `json:"dough"`
	Size string `json:"size"`
	Image string `json:"image"`
	Weight int `json:"weight"`
	Price  int `json:"price"`
	Toppings []*Topping `json:"toppings"`
}

type ProductWithToppings interface {
	ToppingToMask([]int) int
	MaskToToppings(int) []int
}

func (pv *PizzaVariant) ToppingsToMask(toppingIDs []int) int {
	var mask int = 0
	for _, id := range toppingIDs {
			if id > 0 && id <= 64 {
					mask |= 1 << (id - 1)
			}
	}
	return mask
}

func (pv *PizzaVariant) MaskToToppings(mask int) []int {
	var toppings []int
	for i := 0; i < 64; i++ {
			if mask&(1<<i) != 0 {
					toppings = append(toppings, i+1)
			}
	}
	return toppings
}

func (pv *ProductVariant) ToppingsToMask(toppingIDs []int) int {
	var mask int = 0
	for _, id := range toppingIDs {
			if id > 0 && id <= 64 {
					mask |= 1 << (id - 1)
			}
	}
	return mask
}

func (pv *ProductVariant) MaskToToppings(mask int) []int {
	var toppings []int
	for i := 0; i < 64; i++ {
			if mask&(1<<i) != 0 {
					toppings = append(toppings, i+1)
			}
	}
	return toppings
}

type ProductVariant struct {
	ProductId int `json:"productId"`
	Title string `json:"title"`
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