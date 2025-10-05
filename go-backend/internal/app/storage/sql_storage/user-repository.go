package sql_storage

import (
	"database/sql"

	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/storage"
)

type UserRepository struct {
	storage *SQLStorage
}

func (ur *UserRepository) CreateUser(u *model.User) (*model.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}

	if err := ur.storage.db.QueryRow("INSERT INTO users (email, encrypted_password, name, phone) VALUES ($1, $2, $3, $4) RETURNING id", u.Email, u.EncryptedPassword, u.Name, u.Phone).Scan(&u.Id); err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	u.Cart = &model.Cart{}

	if err := ur.storage.db.QueryRow("SELECT id, email, encrypted_password, name, phone FROM users WHERE email = $1", email).Scan(&u.Id, &u.Email, &u.EncryptedPassword, &u.Name, &u.Phone); err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.ErrRecordNotFound
		}
		
		return nil, err
	}

	return u, nil
}

func (ur *UserRepository) FindById(id int) (*model.User, error) {
	u := &model.User{}
	u.Cart = &model.Cart{}

	if err := ur.storage.db.QueryRow("SELECT id, email, encrypted_password, name, phone FROM users WHERE id = $1", id).Scan(&u.Id, &u.Email, &u.EncryptedPassword, &u.Name, &u.Phone); err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.ErrRecordNotFound
		}
		
		return nil, err
	}

	return u, nil
}
