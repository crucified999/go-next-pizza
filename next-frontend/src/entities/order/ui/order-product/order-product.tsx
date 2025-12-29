import { PizzaInCart, ProductInCart } from "@/entities/cart/model/types";
import { convertDough } from "@/entities/product/lib/utils";
import React from "react";

type OrderProductProps = {
  product: ProductInCart;
};

type OrderPizzaProps = {
  pizza: PizzaInCart
}

export const OrderProduct: React.FC<OrderProductProps> = ({ product }) => {
  return (
    <div className="grid grid-cols-[1fr_auto]">
      <div className="flex flex-col gap-2">
        <h4 className="font-bold text-lg">{product.product.title}</h4>
        <span>{product.product.size}, {product.product.weight} г</span>
      </div>

      <div className="flex flex-col gap-1">
        <span>{product.count} x</span>
        <span>{product.count * product.product.price} ₽</span>
      </div>
    </div>
  );
};

export const OrderPizza: React.FC<OrderPizzaProps> = ({ pizza }) => {
  return (
    <div className="grid grid-cols-[1fr_auto]">
      <div className="flex flex-col">
        <h4 className="font-bold text-lg">{pizza.pizza.title}</h4>
        <span className="text-sm text-black/50">{pizza.pizza.size}, {convertDough(pizza.pizza.dough)}, {pizza.pizza.weight} г</span>
      </div>

      <div className="flex flex-col gap-1">
        <span className="self-end font-bold">{pizza.count} x</span>
        <span className="font-bold">{pizza.count * pizza.pizza.price} ₽</span>
      </div>
    </div>
  );
}
