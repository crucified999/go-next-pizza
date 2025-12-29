import { PizzaVariant, ProductVariant } from "@/entities/product/model";
import { Cart } from "../model/types";

const BASE_URL = process.env.NEXT_PUBLIC_API_URL + "/cart";

type ProductInCartRequest = {
  productId: number;
  amount: string;
};

interface DeleteProductRequest extends ProductInCartRequest {
  action: string;
}

type PizzaInCartRequest = {
  pizzaId: number;
  dough: number;
  size: string;
  toppingsMask: number;
};

interface DeletePizzaRequest extends PizzaInCartRequest {
  action: string;
}

export const getCart = async (): Promise<Cart> => {
  return await fetch(`${BASE_URL}`, {
    credentials: "include",
  }).then((res) => res.json());
};

export const addProduct = async ({
  productId,
  amount,
}: ProductInCartRequest): Promise<ProductVariant> => {
  return await fetch(`${BASE_URL}/add-product`, {
    method: "POST",
    credentials: "include",
    body: JSON.stringify({
      productId,
      amount,
    }),
  }).then((res) => res.json());
};

export const deleteProduct = async ({
  productId,
  amount,
  action,
}: DeleteProductRequest) => {
  return await fetch(`${BASE_URL}/delete-product`, {
    method: "DELETE",
    credentials: "include",
    body: JSON.stringify({
      productId,
      amount,
      action,
    }),
  });
};

export const addPizza = async ({
  pizzaId,
  dough,
  size,
  toppingsMask,
}: PizzaInCartRequest): Promise<PizzaVariant> => {
  return await fetch(`${BASE_URL}/add-pizza`, {
    method: "POST",
    credentials: "include",
    body: JSON.stringify({
      pizzaId,
      dough,
      size,
      toppingsMask,
    }),
  }).then((res) => res.json());
};

export const deletePizza = async({
  pizzaId,
  dough,
  size,
  toppingsMask,
  action
}: DeletePizzaRequest) => {
  return await fetch(`${BASE_URL}/delete-pizza`, {
    method: "DELETE",
    credentials: "include",
    body: JSON.stringify({
      pizzaId,
      dough,
      size,
      toppingsMask,
      action
    })
  })
}


export const refreshCart = async () => {
  return await fetch(`${BASE_URL}/refresh`, {
    method: "PUT",
    credentials: "include",
  }) 
}