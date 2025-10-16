package model

type Cart struct {
	Id int `json:"id"`
	UserId int `json:"userId"`
	Products []*Product `json:"products"`
	Combos []*Combo `json:"combos"`
}

type CartProduct struct {
	Product *Product `json:"product"`
	Amount  int      `json:"amount"`
}

type CartCombo struct {
	Combo   *Combo `json:"combo"`
	Amount  int    `json:"amount"`
}