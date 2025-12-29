import { PizzaVariant, ProductVariant } from "@/entities/product/model";
import React from "react";
import { OrderProduct } from "../order-product";
import { Cart, PizzaInCart, ProductInCart } from "@/entities/cart/model/types";
import { OrderPizza } from "../order-product/order-product";
import { declineProductWord } from "@/entities/cart/lib/utils";

type OrderProps = {
  cart: Cart;
};

export const Order: React.FC<OrderProps> = ({ cart }) => {
  return (
    <div className="flex flex-col bg-white shadow-2xl rounded-lg sticky top-25 right-10 p-8 dark:bg-[#101113] dark:border-1 dark:border-white">
      <h3 className="font-bold text-xl mb-8">Состав заказа</h3>
      <ul className="pb-4 flex flex-col gap-4">
        {cart.products.map((p) => (
          <OrderProduct key={p.product.title} product={p} />
        ))}
        {cart.pizzas.map((p) => (
          <OrderPizza key={p.pizza.title} pizza={p} />
        ))}
      </ul>
      <hr />
      <div className="flex flex-col gap-2 py-4">
        <div className="flex justify-between">
          <span className="text-sm font-bold">
            {cart.totalCount} {declineProductWord(cart.totalCount)}
          </span>
          <span className="text-sm font-bold">{cart.totalPrice} ₽</span>
        </div>
        <div className="flex justify-between">
          <span className="text-sm font-bold">Доставка</span>
          <span className="text-sm font-bold">Бесплатно</span>
        </div>
      </div>
      <hr />
      <div className="flex justify-between py-4">
        <span className="font-bold text-[16px]">Сумма заказа</span>
        <span className="font-bold text-[16px]">{cart.totalPrice} ₽</span>
      </div>
    </div>
  );
};
