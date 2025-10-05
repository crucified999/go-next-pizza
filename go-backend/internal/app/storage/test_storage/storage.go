package test_storage

// import (
// 	"github.com/go-next-pizza/internal/app/model"
// 	"github.com/go-next-pizza/internal/app/storage"
// )

// type SQLStorage struct {
// 	userRepository *UserRepository
// 	cartRepository *CartRepository
// }

// func NewSQLStorage() *SQLStorage {
// 	return &SQLStorage{}
// }

// func (s *SQLStorage) User() storage.UserRepository {
// 	if s.userRepository != nil {
// 		return s.userRepository
// 	}

// 	s.userRepository = &UserRepository{
// 		storage: s,
// 		users: make(map[int]*model.User),
// 	}

// 	return s.userRepository
// }

// func (s *SQLStorage) Cart() storage.CartRepository {
// 	if s.cartRepository != nil {
// 		return s.cartRepository
// 	}

// 	s.cartRepository = &CartRepository{
// 		storage: s,
// 		carts: make(map[int]*model.Cart),
// 	}

// 	return s.cartRepository
// }