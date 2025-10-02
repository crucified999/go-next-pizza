package storage

import "github.com/go-next-pizza/internal/app/model"

type UserRepository interface {
	CreateUser(*model.User) error
	FindByEmail(string) (*model.User, error)
	FindById(int) (*model.User, error)
}