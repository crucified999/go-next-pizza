package model

import (
	"fmt"
	"testing"
	"time"
)

func TestUser(t *testing.T) *User {
	return &User{
		Email: fmt.Sprintf("test_email_%d@test.com", time.Now().UnixNano()),
		Password: "2007Coolo_ww!",
		Name: "test_name",
		Phone: "+79155182245",
		Cart: TestCart(t),
		Orders: make([]*Order, 0),
	}
}

func TestCart(t *testing.T) *Cart {
	return &Cart{
		Products: []*Product{},
		Combos: []*Combo{},
	}
}

func TestOrder(userId int, t *testing.T) *Order {
	return &Order{
		UserId: userId,
		Status: "Передаем в ресторан",
		TotalPrice: 100,
		Products: make(map[int]int),
		Combos: make(map[int]int),
		PaymentMethod: "card",
		DeliveryAddress: "test_address",
		DeliveryTime: time.Now(),
		CreatedAt: time.Now(),	
	}
}