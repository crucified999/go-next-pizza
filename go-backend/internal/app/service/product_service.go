package service

import (
	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/storage"
)


type ProductService struct {
	productRepo storage.ProductRepository
}

func NewProductService(productRepo storage.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

func (ps *ProductService) GetProducts() ([]*model.Product, error) {
	products, err := ps.productRepo.GetProducts()

	if err != nil {
		return nil, err
	}

	for index, product := range products {
		product.Category, err = ps.productRepo.GetProductCategory(product.Id)

		if err != nil {
			return nil, err
		}

		variants, err := ps.productRepo.GetProductsVariants(product.Id)

		if err != nil {
			return nil, err
		}

		products[index].Variants = variants

		toppings, err := ps.productRepo.GetProductToppings(product.Id)

		if err != nil {
			return nil, err
		}

		products[index].Toppings = toppings

		if product.Category == "pizza" {
			product.Ingredients, err = ps.productRepo.GetProductIngredients(product.Id)

			if err != nil {
				return nil, err
			}
		}
	}


	return products, nil
}

func (ps *ProductService) GetProductById(id int) (*model.Product, error) {
	product, err := ps.productRepo.GetProductById(id)

	if err != nil {
		return nil, err
	}

	product.Category, err = ps.productRepo.GetProductCategory(product.Id)

	if err != nil {
		return nil, err
	}

	product.Variants, err = ps.productRepo.GetProductsVariants(product.Id)

	if err != nil {
		return nil, err
	}

	product.Ingredients, err = ps.productRepo.GetProductIngredients(product.Id)

	if err != nil {
		return nil, err
	}

	product.Toppings, err = ps.productRepo.GetProductToppings(product.Id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (ps *ProductService) GetProductsByCategory(category string) ([]*model.Product, error) {
	products, err := ps.productRepo.GetProductsByCategory(ps.productRepo.ConvertCategorToId(category))

	if err != nil {
		return nil, err
	}

	for _, product := range products {
		product.Category, err = ps.productRepo.GetProductCategory(product.Id)

		if err != nil {
			return nil, err
		}

		product.Variants, err = ps.productRepo.GetProductsVariants(product.Id)

		if err != nil {
			return nil, err
		}

	}

	return products, nil
}

func (ps *ProductService) GetProductVariant(productId int, size string) (*model.ProductVariant, error) {
	productVariant, err := ps.productRepo.GetProductVariant(productId, size)

	if err != nil {
		return nil, err
	}

	return productVariant, nil
}

func (ps *ProductService) GetPizzaVariant(pizza *model.PizzaVariant) (*model.PizzaVariant, error) {
	pizzaVariant, err := ps.productRepo.GetPizzaVariant(pizza)

	if err != nil {
		return nil, err
	}

	return pizzaVariant, err
}

func (ps *ProductService) GetProductToppings(productId int) ([]*model.Topping, error) {
	return ps.productRepo.GetProductToppings(productId)
}