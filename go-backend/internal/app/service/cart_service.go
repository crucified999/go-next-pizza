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

type ProductInCart struct {
	ProductId int
	Amount string
}

type CartWithItems struct {
	ID         int                   `json:"id"`
	UserID     int                  `json:"userId"`
	Products   []*model.CartProduct `json:"products"`
	Pizzas     []*model.CartPizza   `json:"pizzas"` 
	Combos     []*model.CartCombo   `json:"combos"`
	TotalPrice int                   `json:"totalPrice"`
	TotalCount int                  `json:"totalCount"`
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

	pizzas, err := cs.cartRepo.GetCartPizzas(cart.Id)

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
		Pizzas: pizzas,
		Combos:   combos,
		TotalCount: 0,
	}

	// var prs []*model.ProductVariant

	// for _, p := range cartWithItems.Products {
	// 	prs = append(prs, p.Product)
	// }

	cartWithItems.TotalPrice = cs.calculateTotal(products, combos, pizzas)

	for _, p := range cartWithItems.Products {
		cartWithItems.TotalCount += p.Count
	}

	for _, p := range cartWithItems.Pizzas {
		cartWithItems.TotalCount += p.Count
	}

	return cartWithItems, nil
}

func (cs *CartService) GetCartToppings(mask int) ([]*model.Topping, error) {
	return cs.cartRepo.GetCartToppings(mask)
}

func (cs *CartService) AddProduct(userId int, productId int, amount string) error {
	cart, err := cs.cartRepo.GetCartByUserId(userId)
	if err != nil {
		return err
	}

	if err := cs.cartRepo.AddProduct(productId, amount, cart.Id); err != nil {
		return err
	}

	return nil
}

func (cs *CartService) AddPizza(userId int, pizza *model.PizzaVariant, mask int) error {
	cart, err := cs.cartRepo.GetCartByUserId(userId)

	if err != nil {
		return err
	}

	if err := cs.cartRepo.AddPizza(pizza, mask, cart.Id); err != nil {
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

func (cs *CartService) DeleteProduct(userId int, productId int, amount string) error {
	cart, err := cs.cartRepo.GetCartByUserId(userId)

	if err != nil {
		return err
	}

	if err := cs.cartRepo.DeleteProduct(productId, cart.Id, amount); err != nil {
		return err
	}

	return nil
}

func (cs *CartService) DeletePizza(userId int, pizza *model.PizzaVariant, mask int) error {
	cart, err := cs.cartRepo.GetCartByUserId(userId)

	if err != nil {
		return err
	}

	if err := cs.cartRepo.DeletePizza(pizza, mask, cart.Id); err != nil {
		return err
	}

	return nil
}

func (cs *CartService) DeleteProductCompletely(userId int, productId int, amount string) error {
	cart, err := cs.cartRepo.GetCartByUserId(userId)

	if err != nil {
		return err
	}

	if err := cs.cartRepo.DeleteProductCompletely(productId, cart.Id, amount); err != nil {
		return err
	}

	return nil
}

func (cs *CartService) DeletePizzaCompletely(userId int, pizza *model.PizzaVariant, mask int) error {
	cart, err := cs.cartRepo.GetCartByUserId(userId)

	if err != nil {
		return err
	}

	if err := cs.cartRepo.DeletePizzaCompletely(pizza, mask, cart.Id); err != nil {
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

func (cs *CartService) calculateTotal(products []*model.CartProduct, combos []*model.CartCombo, pizzas []*model.CartPizza) int {
	total := 0

	for _, cartProduct := range products {
		total += cartProduct.Product.Price * cartProduct.Count
	}

	for _, cartPizza := range pizzas {
		total += cartPizza.Pizza.Price * cartPizza.Count
		
		for _, t := range cartPizza.Pizza.Toppings {
			total += t.Price
		}

	}

	for _, cartCombo := range combos {
		total += cartCombo.Combo.Price
	}

	return total
}