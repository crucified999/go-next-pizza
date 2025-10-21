import { Product } from "../model/types";

const BASE_URL = process.env.NEXT_PUBLIC_API_URL;

export const getProducts = (): Promise<Product[]> => {
  return fetch(`${BASE_URL}/products`).then((res) => res.json());
};

