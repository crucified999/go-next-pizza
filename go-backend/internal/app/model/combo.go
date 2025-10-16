package model

type Combo struct {
	Id          int             `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Price       int             `json:"price"`
	Image       string          `json:"image"`
	DefaultProducts []*Product  `json:"defaultProducts"`
	Products    []*Product      `json:"products"`
}