package test_storage_test

// import (
// 	"testing"

// 	"github.com/go-next-pizza/internal/app/model"
// 	"github.com/go-next-pizza/internal/app/storage/test_storage"
// 	"github.com/stretchr/testify/assert"
// )

// func TestCartRepository_Create(t *testing.T) {
// 	s := test_storage.NewSQLStorage()
// 	c := model.TestCart(t)

// 	_, err := s.Cart().CreateCart(c)

// 	assert.NoError(t, err)
// 	assert.NotNil(t, c)
// 	assert.NotNil(t, c.Products)
// 	assert.NotNil(t, c.Combos)
// 	assert.Equal(t, c.Id, 1)
// }

// func TestCartRepository_AddProduct(t *testing.T) {
// 	s := test_storage.NewSQLStorage()
// 	c := model.TestCart(t)

// 	productId := 1

// 	_, err := s.Cart().CreateCart(c)

// 	s.Cart().AddProduct(productId, c)

// 	assert.NoError(t, err)
// 	assert.Equal(t, c.Products[productId], 1)
// }

// func TestCartRepository_AddCombo(t *testing.T) {
// 	s := test_storage.NewSQLStorage()
// 	c := model.TestCart(t)

// 	comboId1 := 1
// 	comboId2 := 2

// 	_, err := s.Cart().CreateCart(c)

// 	s.Cart().AddCombo(comboId1, c)
// 	s.Cart().AddCombo(comboId1, c)
// 	s.Cart().AddCombo(comboId2, c)

// 	assert.NoError(t, err)
// 	assert.Equal(t, c.Combos[comboId1], 2)
// 	assert.Equal(t, c.Combos[comboId2], 1)
// }

// func TestCartRepository_DeleteProduct(t *testing.T) {
// 	s := test_storage.NewSQLStorage()
// 	c := model.TestCart(t)
 
// 	_, err := s.Cart().CreateCart(c)

// 	productId := 1

// 	s.Cart().AddProduct(productId, c)
// 	s.Cart().DeleteProduct(productId, c)

// 	assert.Equal(t, c.Products[productId], 0)

// 	s.Cart().AddProduct(productId, c)
// 	s.Cart().AddProduct(productId, c)
// 	s.Cart().DeleteProduct(productId, c)

// 	assert.NoError(t, err)
// 	assert.Equal(t, c.Products[productId], 1)
// }

// func TestCartRepository_DeleteCombo(t *testing.T) {
// 	s := test_storage.NewSQLStorage()
// 	c := model.TestCart(t)
 
// 	_, err := s.Cart().CreateCart(c)

// 	comboId := 1

// 	s.Cart().AddCombo(comboId, c)
// 	s.Cart().DeleteCombo(comboId, c)

// 	assert.Equal(t, c.Combos[comboId], 0)

// 	s.Cart().AddCombo(comboId, c)
// 	s.Cart().AddCombo(comboId, c)
// 	s.Cart().DeleteCombo(comboId, c)

// 	assert.NoError(t, err)
// 	assert.Equal(t, c.Combos[comboId], 1)
// }

// func TestCartRepository_Refresh(t *testing.T) {
// 	s := test_storage.NewSQLStorage()
// 	c := model.TestCart(t)
	
// 	_, err := s.Cart().CreateCart(c)

// 	comboId := 1
// 	productId := 1

// 	s.Cart().AddProduct(productId, c)
// 	s.Cart().AddCombo(comboId, c)

// 	s.Cart().Refresh(c)

// 	assert.NoError(t, err)
// 	assert.Equal(t, c.Products, map[int]int{})
// 	assert.Equal(t, c.Combos, map[int]int{})
// }