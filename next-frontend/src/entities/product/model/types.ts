export type Product = {
  id: number;
  title: string;
  description: string;
  price: number;
  image: string;
  category: string;
  amount: number;
  weight: number;
  variants?: Variant[];
  ingredients?: PizzaIngredient[];
  toppings?: Topping[];
};

export type Pizza = Product & {
  ingredients: PizzaIngredient[];
}

export type Variant = {
  productId: number;
  doughType?: number;
  price: number;
  size: string;
  image: string;
  weight: number;
}

export type PizzaIngredient = {
  id: number;
  title: string;
  replacable: boolean;
}

export type Topping = {
  id: number;
  title: string;
  price: number;
  image: string;
}

// type CustomPizzaIngredient struct {
// 	ID           int     `json:"id"`
// 	CustomPizzaID int     `json:"custom_pizza_id"`
// 	IngredientID int     `json:"ingredient_id"`
// 	IsAdded      bool    `json:"is_added"`
// 	CreatedAt    time.Time `json:"created_at"`
// }

export type CustomPizzaIngredient = {
  id: number;
  customPizzaId: number;
  ingredientId: number;
  isAdded: boolean;
}

export type CustomPizza = {
  id: number;
  totalPrice: number;
  ingredients: CustomPizzaIngredient[];
  dough: string;
  size: string;
}

// type CustomPizza struct {
// 	ID          int                      `json:"id"`
// 	UserID      int                      `json:"user_id"`
// 	BasePizzaID int                      `json:"base_pizza_id"`
// 	Name        string                   `json:"name"`
// 	TotalPrice  float64                  `json:"total_price"`
// 	Ingredients []CustomPizzaIngredient  `json:"ingredients"`
// 	Dough				string									 `json:"dough"`
// 	Size 				string									 `json:"size"`
// 	CreatedAt   time.Time                `json:"created_at"`
// 	UpdatedAt   time.Time                `json:"updated_at"`
// }

export type CreateCustomPizzaRequest = {
  base_pizza_id: number;
  name: string;
  ingredients: CustomPizzaIngredientRequest[];
  dough: string;
  size: string;
}

export type CreateCustomPizzaResponse = {
  id: number,

}

export type CustomPizzaIngredientRequest = {
  ingredientId: number;
  action: string; 
}

export type CustomPizzaUpdateRequest = {
  id: number;
  ingredients?: CustomPizzaIngredient[];
  dough?: string;
  size?: string;
}