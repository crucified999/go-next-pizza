package storage

type Storage interface {
	User() UserRepository
	Cart() CartRepository
	Order() OrderRepository
	Product() ProductRepository
	Ingredient() IngredientRepository
	CustomPizza() CustomPizzaRepository
}