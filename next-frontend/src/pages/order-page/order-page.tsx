"use client";

import { useAppDispatch, useAppSelector } from "@/app/store";
import { refreshCart } from "@/entities/cart/lib/api";
import { clearCart, fetchCart } from "@/entities/cart/store/cartSlice";
import { createOrder } from "@/entities/order/lib/api";
import { CardForm } from "@/entities/order/ui/card-form";
import { Order } from "@/entities/order/ui/order";
import { OrderForm } from "@/entities/order/ui/order-form/order-form";
import { checkUserAuth } from "@/entities/user/store/userSlice";
import { CheckoutContainer } from "@/features/checkout/ui";
import { Footer } from "@/shared/ui/footer";
import { Logo } from "@/shared/ui/logo";
import { useEffect } from "react";

export const OrderPage = () => {
  const dispatch = useAppDispatch();
  const user = useAppSelector((state) => state.user);
  const cart = useAppSelector((state) => state.cart.cart);

  const handleSubmit = async () => {
    await createOrder({ totalPrice: cart.totalPrice })
    await refreshCart();

    dispatch(clearCart());
  }
  

  useEffect(() => {
    document.body.style.overflow = "auto";

    dispatch(checkUserAuth());
    dispatch(fetchCart());
  }, [dispatch]);

  return (
    <>
      <div className="grid grid-cols-2">
        <div>
          <div className="grid grid-cols-2">
            <Logo />
          </div>
          <div>
            <h1 className="font-bold text-4xl mt-15">Заказ на доставку</h1>
            <OrderForm user={user} />
            <CardForm className="my-10" amount={cart.totalPrice} onSubmit={handleSubmit} />
          </div>
        </div>
        <div className="flex flex-col gap-20">
          <CheckoutContainer />
          <Order cart={cart} />
        </div>
      </div>
      <Footer />
    </>
  );
};
