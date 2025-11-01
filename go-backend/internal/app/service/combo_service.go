package service

import (
	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/storage"
)


type ComboService struct {
	comboRepo storage.ComboRepository
	productRepo storage.ProductRepository
}

func NewComboService(comboRepo storage.ComboRepository, productRepo storage.ProductRepository) *ComboService {
	return &ComboService{
		comboRepo: comboRepo,
		productRepo: productRepo,
	}
}

func (cs *ComboService) GetCombos() ([]*model.Combo, error) {
	combos, err := cs.comboRepo.GetCombos()
	if err != nil {
		return nil, err
	}

	for index, combo := range combos {
		products, err := cs.comboRepo.GetComboProducts(combo.Id)
		if err != nil {
			return nil, err
		}

		for _, product := range products {
			product.Category, err = cs.productRepo.GetProductCategory(product.Id)

			if err != nil {
				return nil, err
			}

			product.Variants, err = cs.productRepo.GetProductsVariants(product.Id)
			if err != nil {
				return nil, err
			}

			product.Ingredients, err = cs.productRepo.GetProductIngredients(product.Id)
			if err != nil {
				return nil, err
			}

			product.Toppings, err = cs.productRepo.GetProductToppings(product.Id)
			if err != nil {
				return nil, err
			}
		}

		defaultProducts, err := cs.comboRepo.GetComboDefaultProducts(combo.Id)

		if err != nil {
			return nil, err
		}

		for _, product := range defaultProducts {
			product.Category, err = cs.productRepo.GetProductCategory(product.Id)
			if err != nil {
				return nil, err
			}

			variants, err := cs.productRepo.GetProductsVariants(product.Id)
			if err != nil {
				return nil, err
			}

			if variants != nil {
				product.Variants = append(product.Variants, variants[2])
			}

			product.Ingredients, err = cs.productRepo.GetProductIngredients(product.Id)
			if err != nil {
				return nil, err
			}

			product.Toppings, err = cs.productRepo.GetProductToppings(product.Id)
			if err != nil {
				return nil, err
			}
		}

		combos[index].DefaultProducts = defaultProducts
		combos[index].Products = products
	}

	return combos, nil
}

func (cs *ComboService) GetComboById(id int) (*model.Combo, error) {
	combo, err := cs.comboRepo.GetComboById(id)
	if err != nil {
		return nil, err
	}

	products, err := cs.comboRepo.GetComboProducts(combo.Id)

	for _, product := range products {
			product.Category, err = cs.productRepo.GetProductCategory(product.Id)
			if err != nil {
				return nil, err
			}

			variants, err := cs.productRepo.GetProductsVariants(product.Id)
			if err != nil {
				return nil, err
			}

			if variants != nil {
				product.Variants = append(product.Variants, variants[2])
			}

			product.Ingredients, err = cs.productRepo.GetProductIngredients(product.Id)
			if err != nil {
				return nil, err
			}

			product.Toppings, err = cs.productRepo.GetProductToppings(product.Id)
			if err != nil {
				return nil, err
			}
		}

		combo.Products = products

		defaultProducts, err := cs.comboRepo.GetComboDefaultProducts(combo.Id)
		if err != nil {
			return nil, err
		}

		for _, product := range defaultProducts {
			product.Category, err = cs.productRepo.GetProductCategory(product.Id)
			if err != nil {
				return nil, err
			}

			variants, err := cs.productRepo.GetProductsVariants(product.Id)
			if err != nil {
				return nil, err
			}

			if variants != nil {
				product.Variants = append(product.Variants, variants[2])
			}

			product.Ingredients, err = cs.productRepo.GetProductIngredients(product.Id)
			if err != nil {
				return nil, err
			}

			product.Toppings, err = cs.productRepo.GetProductToppings(product.Id)
			if err != nil {
				return nil, err
			}
		}

		combo.DefaultProducts = defaultProducts

		return combo, nil
}

func (cs *ComboService) ReplaceProduct(productToReplaceId int, replacerId int, combo *model.Combo) (*model.Combo, error) {
	product, err := cs.productRepo.GetProductById(replacerId)
	if err != nil {
		return nil, err
	}

	for i := len(combo.Products) - 1; i >= 0; i-- {
		if combo.Products[i].Id == productToReplaceId {
			combo.Products = append(combo.Products[:i], combo.Products[i+1:]...)
			break
		}
	}

	combo.Products = append(combo.Products, product)

	return combo, nil
}	