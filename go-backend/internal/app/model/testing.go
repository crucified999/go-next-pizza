package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email: "test_email@test.com",
		Password: "2007Coolo_ww!",
	}
}