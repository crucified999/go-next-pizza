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