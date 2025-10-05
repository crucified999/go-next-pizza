package test_storage

// import (
// 	"github.com/go-next-pizza/internal/app/model"
// )

// type CartRepository struct {
// 	storage *SQLStorage
// 	carts map[int]*model.Cart
// }

// func (cr *CartRepository) CreateCart(c *model.Cart) (*model.Cart, error) {

// 	c.Id = len(cr.carts) + 1
// 	c.Products = make(map[int]int)
// 	cr.carts[c.Id] = c


// 	return c, nil

// } 

// func (cr *CartRepository) AddProduct(productId int, c *model.Cart) error {
// 	_, ok := cr.carts[c.Id].Products[productId]

// 	if !ok {
// 		cr.carts[c.Id].Products[productId] = 1
// 		return nil
// 	}

// 	cr.carts[c.Id].Products[productId] += 1

// 	return nil
// }

// func (cr *CartRepository) AddCombo(comboId int, c *model.Cart) error {
// 	_, ok := cr.carts[c.Id].Combos[comboId]

// 	if !ok {
// 		cr.carts[c.Id].Combos[comboId] = 1
// 		return nil
// 	}

// 	cr.carts[c.Id].Combos[comboId] += 1

// 	return nil
// }

// func (cr *CartRepository) DeleteProduct(productId int, c *model.Cart) error {
// 	if _, exists := cr.carts[c.Id].Products[productId]; !exists {
// 		return nil
// 	}

// 	cr.carts[c.Id].Products[productId] -= 1

// 	if cr.carts[c.Id].Products[productId] <= 0 {
// 		delete(cr.carts[c.Id].Products, productId)
// 	}

// 	return nil
// }

// func (cr *CartRepository) DeleteCombo(comboId int, c *model.Cart) error {
// 	if _, exists := cr.carts[c.Id].Combos[comboId]; !exists {
// 		return nil
// 	}

// 	cr.carts[c.Id].Combos[comboId] -= 1

// 	if cr.carts[c.Id].Combos[comboId] == 0 {
// 		delete(cr.carts[c.Id].Combos, comboId)
// 	}

// 	return nil
// }

// func (cr *CartRepository) Refresh(c *model.Cart) error {
// 	for p := range cr.carts[c.Id].Products {
// 		delete(cr.carts[c.Id].Products, p)
// 	}

// 	for combo := range cr.carts[c.Id].Combos {
// 		delete(cr.carts[c.Id].Combos, combo)
// 	}

// 	return nil
// }