package service

import (
	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/storage"
)

type CartService struct {
	cartRepo storage.CartRepository
	productRepo storage.ProductRepository
	comboRepo storage.ComboRepository
}

type CartWithItems struct {
	ID         int                 `json:"id"`
	UserID     int                 `json:"userId"`
	Products   []*model.CartProduct `json:"products"`
	Combos     []*model.CartCombo   `json:"combos"`
	TotalPrice float64             `json:"totalPrice"`
}

func NewCartService(
	cartRepo storage.CartRepository, 
	productRepo storage.ProductRepository, 
	comboRepo storage.ComboRepository,
	) *CartService {
		return &CartService{
			cartRepo: cartRepo,
			productRepo: productRepo,
			comboRepo: comboRepo,
		}
}

func (cs *CartService) GetCartByUserId(userId int) (*CartWithItems, error) {
	cart, err := cs.cartRepo.GetCartByUserId(userId)

	if err != nil {
		return nil, err
	}

	products, err := cs.cartRepo.GetCartProducts(cart.Id)

	if err != nil {
		return nil, err
	}

	combos, err := cs.cartRepo.GetCartCombos(cart.Id)

	if err != nil {
		return nil, err
	}
	
	cartWithItems := &CartWithItems{
		ID:       cart.Id,
		UserID:   cart.UserId,
		Products: products,
		Combos:   combos,
	}

	cartWithItems.TotalPrice = cs.calculateTotal(products, combos)

	return cartWithItems, nil
}

func (cs *CartService) AddProduct(userId int, productId int) error {
	cart, err := cs.cartRepo.GetCartByUserId(userId)
	if err != nil {
		return err
	}


	if err := cs.cartRepo.AddProduct(productId, cart.Id); err != nil {
		return err
	}


	return nil
}

func (cs *CartService) AddCombo(userId int, comboId int) error {
	cart, err := cs.cartRepo.GetCartByUserId(userId)

	if err != nil {
		return err
	}

	err = cs.cartRepo.AddCombo(comboId, cart.Id)

	if err != nil {
		return err
	}
	
	return nil
}

func (cs *CartService) DeleteProduct(productId int, cartId int) error {
	err := cs.cartRepo.DeleteProduct(productId, cartId)

	if err != nil {
		return err
	}

	return nil
}

func (cs *CartService) DeleteCombo(comboId int, cartId int) error {
	err := cs.cartRepo.DeleteCombo(comboId, cartId)

	if err != nil {
		return err
	}

	return nil
}

func (cs *CartService) Refresh(cartId int) error {
	err := cs.cartRepo.Refresh(cartId)

	if err != nil {
		return err
	}

	return nil
}

func (cs *CartService) calculateTotal(products []*model.CartProduct, combos []*model.CartCombo) float64 {
	total := 0.0

	for _, cartProduct := range products {
		if cartProduct.Product != nil && cartProduct.Product.Price.Valid {
			productPrice := float64(cartProduct.Product.Price.Int64)
			total += productPrice * float64(cartProduct.Amount)
		}
	}

	for _, cartCombo := range combos {
		if cartCombo.Combo != nil && cartCombo.Combo.Price != 0 {
			comboPrice := float64(cartCombo.Combo.Price)
			total += comboPrice * float64(cartCombo.Amount)
		}
	}

	return total
}