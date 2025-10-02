package test_storage

import (
	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/storage"
)

type SQLStorage struct {
	userRepository *UserRepository
}

func NewSQLStorage() *SQLStorage {
	return &SQLStorage{}
}

func (s *SQLStorage) User() storage.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		storage: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}