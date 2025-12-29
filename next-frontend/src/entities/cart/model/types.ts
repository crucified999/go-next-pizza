import { Combo } from "@/entities/combo/model";
import { PizzaVariant, ProductVariant } from "@/entities/product/model"

export type ProductInCart = {
  product: ProductVariant;
  count: number;
}

export type PizzaInCart = {
  pizza: PizzaVariant;
  count: number;
}

export type Cart = {
  products: ProductInCart[];
  pizzas: PizzaInCart[];
  combos: Combo[];
  totalCount: number;
  totalPrice: number;
}