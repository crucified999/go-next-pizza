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
	// Получаем базовую пиццу
	basePizza, err := cps.productRepo.GetProductByID(req.BasePizzaID)
	if err != nil {
		return nil, err
	}

	// Создаем кастомную пиццу
	customPizza := &model.CustomPizza{
		UserID:      userID,
		BasePizzaID: req.BasePizzaID,
		Name:        req.Name,
		TotalPrice:  basePizza.Price,
		Ingredients: []model.CustomPizzaIngredient{},
	}

	// Обрабатываем изменения ингредиентов
	for _, ingredientReq := range req.Ingredients {
		ingredient, err := cps.ingredientRepo.GetIngredientByID(ingredientReq.IngredientID)
		if err != nil {
			return nil, err
		}

		customIngredient := model.CustomPizzaIngredient{
			IngredientID: ingredientReq.IngredientID,
			IsAdded:      ingredientReq.Action == "add",
		}

		// Проверяем, является ли ингредиент стандартным для базовой пиццы
		isStandard := cps.isStandardIngredient(req.BasePizzaID, ingredientReq.IngredientID)

		if ingredientReq.Action == "add" {
			// Добавляем ингредиент
			customPizza.Ingredients = append(customPizza.Ingredients, customIngredient)
			customPizza.TotalPrice += ingredient.Price
		} else if ingredientReq.Action == "remove" {
			// Удаляем стандартный ингредиент
			if isStandard {
				customPizza.Ingredients = append(customPizza.Ingredients, customIngredient)
				customPizza.TotalPrice -= ingredient.Price
			}
		}
	}

	// Сохраняем в базу данных
	return cps.customPizzaRepo.CreateCustomPizza(customPizza)
}

func (cps *CustomPizzaService) GetCustomPizzaByID(id int) (*model.CustomPizza, error) {
	return cps.customPizzaRepo.GetCustomPizzaByID(id)
}

func (cps *CustomPizzaService) GetCustomPizzasByUserID(userID int) ([]*model.CustomPizza, error) {
	return cps.customPizzaRepo.GetCustomPizzasByUserID(userID)
}

func (cps *CustomPizzaService) UpdateCustomPizza(id int, userID int, req model.UpdateCustomPizzaRequest) (*model.CustomPizza, error) {
	// Проверяем, что кастомная пицца принадлежит пользователю
	customPizza, err := cps.customPizzaRepo.GetCustomPizzaByID(id)
	if err != nil {
		return nil, err
	}

	if customPizza.UserID != userID {
		return nil, errors.New("custom pizza doesn't belong to user")
	}

	// Получаем базовую пиццу для пересчета цены
	basePizza, err := cps.productRepo.GetProductByID(customPizza.BasePizzaID)
	if err != nil {
		return nil, err
	}

	// Обновляем информацию
	customPizza.Name = req.Name
	customPizza.TotalPrice = basePizza.Price
	customPizza.Ingredients = []model.CustomPizzaIngredient{}

	// Обрабатываем изменения ингредиентов
	for _, ingredientReq := range req.Ingredients {
		ingredient, err := cps.ingredientRepo.GetIngredientByID(ingredientReq.IngredientID)
		if err != nil {
			return nil, err
		}

		customIngredient := model.CustomPizzaIngredient{
			IngredientID: ingredientReq.IngredientID,
			IsAdded:      ingredientReq.Action == "add",
		}

		// Проверяем, является ли ингредиент стандартным для базовой пиццы
		isStandard := cps.isStandardIngredient(customPizza.BasePizzaID, ingredientReq.IngredientID)

		if ingredientReq.Action == "add" {
			// Добавляем ингредиент
			customPizza.Ingredients = append(customPizza.Ingredients, customIngredient)
			customPizza.TotalPrice += ingredient.Price
		} else if ingredientReq.Action == "remove" {
			// Удаляем стандартный ингредиент
			if isStandard {
				customPizza.Ingredients = append(customPizza.Ingredients, customIngredient)
				customPizza.TotalPrice -= ingredient.Price
			}
		}
	}

	// Сохраняем изменения
	return cps.customPizzaRepo.UpdateCustomPizza(customPizza)
}

func (cps *CustomPizzaService) DeleteCustomPizza(id int, userID int) error {
	// Проверяем, что кастомная пицца принадлежит пользователю
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
	// Здесь должна быть логика проверки, является ли ингредиент стандартным для пиццы
	// Пока возвращаем false, так как у нас нет таблицы связи пицца-ингредиент
	// В реальном проекте здесь был бы запрос к базе данных
	return false
}