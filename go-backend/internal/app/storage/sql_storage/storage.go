package sql_storage

import (
	"database/sql"

	"github.com/go-next-pizza/internal/app/storage"
	_ "github.com/lib/pq"
)

type SQLStorage struct {
	db *sql.DB
	userRepository *UserRepository
}

func NewSQLStorage(db *sql.DB) *SQLStorage {
	return &SQLStorage{
		db: db,
	}
}

func (s *SQLStorage) User() storage.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		storage: s,
	}

	return s.userRepository
}