import { Order } from "../model/types";

const BASE_URL = process.env.NEXT_PUBLIC_API_URL;

type CreateOrderRequest = {
  totalPrice: number;
}

export const createOrder = async ({ totalPrice }: CreateOrderRequest) => {
  return await fetch(`${BASE_URL}/orders`, {
    method: "POST",
    credentials: "include",
    body: JSON.stringify({ totalPrice })
  })
}

export const getOrders = async (): Promise<Order[]> => {
  return await fetch(`${BASE_URL}/users/orders`, {
    method: "GET",
    credentials: "include",
  })
  .then((res) => res.json());
}