package model

type Cart struct {
	Id int `json:"id"`
	UserId int `json:"userId"`
	Products []*CartProduct `json:"products"`
	Combos []*Combo `json:"combos"`
	TotalCount int `json:"total_count"`
}

type CartProduct struct {
	Product *ProductVariant `json:"product"`
	Count  int      `json:"count"`
}

type CartPizza struct {
	Pizza *PizzaVariant `json:"pizza"`
	Count int `json:"count"`
}

type CartCombo struct {
	Combo   *Combo `json:"combo"`
	Amount  int    `json:"amount"`
}