package model

import (
	"testing"
	"time"
)

func TestUser(t *testing.T) *User {
	return &User{
		Phone: "+79155182245",
		Cart: TestCart(t),
		Orders: make([]*Order, 0),
	}
}

func TestCart(t *testing.T) *Cart {
	return &Cart{
		Products: []*CartProduct{},
		Combos: []*Combo{},
	}
}

func TestOrder(userId int, t *testing.T) *Order {
	return &Order{
		UserId: userId,
		TotalPrice: 100,
		CreatedAt: time.Now(),	
	}
}