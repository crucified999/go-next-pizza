package test_storage

// import (
// 	"github.com/go-next-pizza/internal/app/model"
// 	"github.com/go-next-pizza/internal/app/storage"
// )

// type UserRepository struct {
// 	storage *SQLStorage
// 	users   map[int]*model.User
// }


// func (ur *UserRepository) CreateUser(u *model.User) error {
// 	if err := u.Validate(); err != nil {
// 		return err
// 	}

// 	if err := u.BeforeCreate(); err != nil {
// 		return err
// 	}

// 	u.Id = len(ur.users) + 1
// 	ur.users[u.Id] = u

// 	u.Cart, _ = ur.storage.Cart().CreateCart(u.Cart)

// 	return nil
// }

// func (ur *UserRepository) FindByEmail(email string) (*model.User, error) {
	
// 	for _, u := range ur.users {
// 		if u.Email == email {
// 			return u, nil
// 		}
// 	}

// 	return nil, storage.ErrRecordNotFound
// }

// func (ur *UserRepository) FindById(id int) (*model.User, error) {
// 	u, ok := ur.users[id]

// 	if !ok {
// 		return nil, storage.ErrRecordNotFound
// 	}

// 	return u, nil
// }
