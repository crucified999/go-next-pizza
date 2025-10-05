package model

type Cart struct {
	Id int `json:"id"`
	UserId int `json:"userId"`
	Products map[int]int `json:"products"`
	Combos map[int]int `json:"combos"`
}