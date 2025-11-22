import {
  CreateCustomPizzaRequest,
  CustomPizza,
  CustomPizzaUpdateRequest,
  Product,
} from "../model/types";

const BASE_URL = process.env.NEXT_PUBLIC_API_URL;

export const getProducts = (): Promise<Product[]> => {
  return fetch(`${BASE_URL}/products`).then((res) => res.json());
};

export const getProductById = (id: number): Promise<Product> => {
  return fetch(`${BASE_URL}/products/${id}`).then((res) => res.json());
};

export const createCustomPizza = (
  req: CreateCustomPizzaRequest
): Promise<CustomPizza> => {
  return fetch(`${BASE_URL}/custom-pizzas`, {
    method: "POST",
    credentials: "include",
    body: JSON.stringify(req),
  })
    .then((res) => res.json())
    .then((data) => {
      return data;
    });
};

export const updateCustomPizza = (req: CustomPizzaUpdateRequest): Promise<CustomPizza> => {
  return fetch(`${BASE_URL}/custom-pizzas/${req.id}`, {
    method: "PUT",
    credentials: "include",
    body: JSON.stringify(req),
  })
    .then((res) => res.json())
    .then((data) => data);
};

export const deleteCustomPizza = (id: number) => {
  return fetch(`${BASE_URL}/custom-pizzas/${id}`, {
    method: "DELETE",
    credentials: "include",
  });
};

export const getCustomPizza = (id: number): Promise<CustomPizza> => {
  return fetch(`${BASE_URL}/custom-pizzas/${id}`, {
    credentials: "include"
  })
  .then((res) => res.json());
}