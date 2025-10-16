package service

import (
	"errors"

	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/storage"
)

type CustomPizzaService struct {
	customPizzaRepo storage.CustomPizzaRepository
	productRepo     storage.ProductRepository
	ingredientRepo  storage.IngredientRepository
}

func NewCustomPizzaService(
	customPizzaRepo storage.CustomPizzaRepository,
	productRepo storage.ProductRepository,
	ingredientRepo storage.IngredientRepository,
) *CustomPizzaService {
	return &CustomPizzaService{
		customPizzaRepo: customPizzaRepo,
		productRepo:     productRepo,
		ingredientRepo:  ingredientRepo,
	}
}

func (cps *CustomPizzaService) CreateCustomPizza(userID int, req model.CreateCustomPizzaRequest) (*model.CustomPizza, error) {
	basePizza, err := cps.productRepo.GetProductById(req.BasePizzaID)

	if err != nil {
		return nil, err
	}

	var basePrice float64
	if basePizza.Price.Valid {
		basePrice = float64(basePizza.Price.Int64)
	}
	
	customPizza := &model.CustomPizza{
		UserID:      userID,
		BasePizzaID: req.BasePizzaID,
		Name:        req.Name,
		TotalPrice:  basePrice,
		Ingredients: []model.CustomPizzaIngredient{},
	}

	for _, ingredientReq := range req.Ingredients {
		ingredient, err := cps.ingredientRepo.GetIngredientByID(ingredientReq.IngredientID)
		if err != nil {
			return nil, err
		}

		customIngredient := model.CustomPizzaIngredient{
			IngredientID: ingredientReq.IngredientID,
			IsAdded:      ingredientReq.Action == "add",
		}

		isStandard := cps.isStandardIngredient(req.BasePizzaID, ingredientReq.IngredientID)

		if ingredientReq.Action == "add" {
			customPizza.Ingredients = append(customPizza.Ingredients, customIngredient)
			customPizza.TotalPrice += ingredient.Price
		} else if ingredientReq.Action == "remove" {
			if isStandard {
				customPizza.Ingredients = append(customPizza.Ingredients, customIngredient)
				customPizza.TotalPrice -= ingredient.Price
			}
		}
	}

	return cps.customPizzaRepo.CreateCustomPizza(customPizza)
}

func (cps *CustomPizzaService) GetCustomPizzaByID(id int) (*model.CustomPizza, error) {
	return cps.customPizzaRepo.GetCustomPizzaByID(id)
}

func (cps *CustomPizzaService) GetCustomPizzasByUserID(userID int) ([]*model.CustomPizza, error) {
	return cps.customPizzaRepo.GetCustomPizzasByUserID(userID)
}

func (cps *CustomPizzaService) UpdateCustomPizza(id int, userID int, req model.UpdateCustomPizzaRequest) (*model.CustomPizza, error) {
	customPizza, err := cps.customPizzaRepo.GetCustomPizzaByID(id)

	if err != nil {
		return nil, err
	}

	if customPizza.UserID != userID {
		return nil, errors.New("custom pizza doesn't belong to user")
	}

	basePizza, err := cps.productRepo.GetProductById(customPizza.BasePizzaID)
	if err != nil {
		return nil, err
	}

	var basePrice float64
	if basePizza.Price.Valid {
		basePrice = float64(basePizza.Price.Int64)
	}
	
	customPizza.Name = req.Name
	customPizza.TotalPrice = basePrice
	customPizza.Ingredients = []model.CustomPizzaIngredient{}

	for _, ingredientReq := range req.Ingredients {
		ingredient, err := cps.ingredientRepo.GetIngredientByID(ingredientReq.IngredientID)
		if err != nil {
			return nil, err
		}

		customIngredient := model.CustomPizzaIngredient{
			IngredientID: ingredientReq.IngredientID,
			IsAdded:      ingredientReq.Action == "add",
		}

		isStandard := cps.isStandardIngredient(customPizza.BasePizzaID, ingredientReq.IngredientID)

		if ingredientReq.Action == "add" {
			customPizza.Ingredients = append(customPizza.Ingredients, customIngredient)
			customPizza.TotalPrice += ingredient.Price
		} else if ingredientReq.Action == "remove" {
			if isStandard {
				customPizza.Ingredients = append(customPizza.Ingredients, customIngredient)
				customPizza.TotalPrice -= ingredient.Price
			}
		}
	}

	return cps.customPizzaRepo.UpdateCustomPizza(customPizza)
}

func (cps *CustomPizzaService) DeleteCustomPizza(id int, userID int) error {
	customPizza, err := cps.customPizzaRepo.GetCustomPizzaByID(id)

	if err != nil {
		return err
	}

	if customPizza.UserID != userID {
		return errors.New("custom pizza doesn't belong to user")
	}

	return cps.customPizzaRepo.DeleteCustomPizza(id)
}

func (cps *CustomPizzaService) isStandardIngredient(pizzaID int, ingredientID int) bool {
	return false
}