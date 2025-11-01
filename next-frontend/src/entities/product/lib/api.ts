import { Product } from "../model/types";

const BASE_URL = process.env.NEXT_PUBLIC_API_URL;

export const getProducts = (): Promise<Product[]> => {
  return fetch(`${BASE_URL}/products`).then((res) => res.json());
};

export const getProductById = (id: number): Promise<Product> => {
  return fetch(`${BASE_URL}/products/${id}`).then((res) => res.json());
};