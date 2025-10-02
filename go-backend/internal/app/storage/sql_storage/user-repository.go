package sql_storage

import (
	"database/sql"

	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/storage"
)

type UserRepository struct {
	storage *SQLStorage
}

func (ur *UserRepository) CreateUser(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

    if err := ur.storage.db.QueryRow("INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id", u.Email, u.EncryptedPassword).Scan(&u.Id); err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}

	if err := ur.storage.db.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&u.Id, &u.Email, &u.EncryptedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.ErrRecordNotFound
		}
		
		return nil, err
	}

	return u, nil
}

func (ur *UserRepository) FindById(id int) (*model.User, error) {
	u := &model.User{}

    if err := ur.storage.db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&u.Id, &u.Email, &u.EncryptedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.ErrRecordNotFound
		}
		
		return nil, err
	}

	return u, nil
}