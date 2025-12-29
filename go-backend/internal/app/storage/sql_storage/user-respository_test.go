package sql_storage_test

import (
	"testing"

	"github.com/go-next-pizza/internal/app/model"
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
