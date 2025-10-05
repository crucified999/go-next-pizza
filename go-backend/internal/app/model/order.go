package model

import "time"

type Order struct {
	Id int `json:"id"`
	UserId int `json:"userId"`
	Status string `json:"status"`
	TotalPrice int `json:"totalPrice"`
	PaymentMethod string `json:"paymentMethod"`
	DeliveryAddress string `json:"deliveryAddress"`
	Products map[int]int `json:"products"`
	Combos map[int]int `json:"combos"`
	DeliveryTime time.Time `json:"deliveryTime"`
	CreatedAt time.Time `json:"createdAt"`
}

func (o *Order) ChangeStatus(status string) {
	seconds := time.Since(o.CreatedAt).Seconds()
	duration := time.Since(o.CreatedAt)

	if seconds < 30 {
		o.Status = "Передаем в ресторан"
	} else if seconds >= 30 && seconds < 600 {
		o.Status = "Готовим заказ"
	} else if seconds >= 600 && duration < 900 {
		o.Status = "Везем заказ"
	} else if seconds >= 900 {
		o.Status = "Доставлен"
	}
}