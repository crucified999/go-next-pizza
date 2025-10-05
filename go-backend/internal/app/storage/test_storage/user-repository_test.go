package test_storage_test

// import (
// 	"testing"

// 	"github.com/go-next-pizza/internal/app/model"
// 	"github.com/go-next-pizza/internal/app/storage"
// 	"github.com/go-next-pizza/internal/app/storage/test_storage"
// 	"github.com/stretchr/testify/assert"
	
// )

// func TestUserRepository_Create(t *testing.T) {
// 	s := test_storage.NewSQLStorage()
	
// 	u := model.TestUser(t)

// 	assert.NoError(t, s.User().CreateUser(u))
// 	assert.NotNil(t, u)
// 	assert.NotNil(t, u.Cart)
// }

// func TestUserRepository_FindByEmail_IsNotExist(t *testing.T) {
// 	s := test_storage.NewSQLStorage()

// 	exampleEmail := "test_email@test.com"

// 	_, err := s.User().FindByEmail(exampleEmail)

// 	assert.EqualError(t, err, storage.ErrRecordNotFound.Error())
// }

// func TestUserRepository_FindByEmail_IsExist(t *testing.T) {
// 	s := test_storage.NewSQLStorage()

// 	exampleEmail := "test_email@test.com"

// 	u := model.TestUser(t)
// 	u.Email = exampleEmail

//   s.User().CreateUser(u)

// 	u, err := s.User().FindByEmail(exampleEmail)

// 	assert.NoError(t, err)
// 	assert.NotNil(t, u)
// }


// func TestUserRepository_FindById(t *testing.T) {
// 	s := test_storage.NewSQLStorage()
// 	u := model.TestUser(t)

//   s.User().CreateUser(u)

// 	u1, err := s.User().FindById(u.Id)

// 	assert.NoError(t, err)
// 	assert.NotNil(t, u1)
// }