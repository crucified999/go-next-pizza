package sql_storage_test

import (
	"testing"

	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/storage"
	sqlstorage "github.com/go-next-pizza/internal/app/storage/sql_storage"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_CreateUser(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstorage.NewSQLStorage(db)
	u := model.TestUser(t)

	_, err := s.User().CreateUser(u)
	
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.NotNil(t, u.Cart)
}

func TestUserRepository_FindByEmail_IsNotExist(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstorage.NewSQLStorage(db)

	exampleEmail := "test_email@test.com"

	_, err := s.User().FindByEmail(exampleEmail)

	assert.Error(t, err, storage.ErrRecordNotFound.Error())
}

func TestUserRepository_FindByEmail_IsExist(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstorage.NewSQLStorage(db)

	exampleEmail := "test_email@test.com"

	u := model.TestUser(t)
	u.Email = exampleEmail

  s.User().CreateUser(u)

	u, err := s.User().FindByEmail(exampleEmail)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindById(t *testing.T) {
	db, teardown := sqlstorage.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstorage.NewSQLStorage(db)
	u := model.TestUser(t)

  s.User().CreateUser(u)

	u2, err := s.User().FindById(u.Id)

	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
