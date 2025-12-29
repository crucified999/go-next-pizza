'use client';

import { useAppDispatch, useAppSelector } from "@/app/store";
import { Cart } from "@/entities/cart/ui/cart";
import { setCurrentCategory } from "@/entities/category/store/categorySlice";
import { fetchOrders } from "@/entities/order/store/orderSlice";
import { OrderStory } from "@/entities/order/ui/order-story/order-story";
import { checkUserAuth } from "@/entities/user/store/userSlice";
import { LogoutButton } from "@/entities/user/ui/logout-button/logout-button";
import { PersonalDataForm } from "@/entities/user/ui/personal-data-form";
import { Footer } from "@/shared/ui/footer";
import { Header } from "@/shared/ui/header";
import { PostHeader } from "@/shared/ui/post-header";
import { useEffect } from "react";

export const ProfilePage = () => {
  const dispatch = useAppDispatch();
  const orders = useAppSelector((state) => state.order.orders);

  useEffect(() => {
    dispatch(checkUserAuth());
    dispatch(fetchOrders());
    dispatch(setCurrentCategory(""));
  }, [dispatch]);

  return (
    <>
      <Header />
      <Cart />
      <PostHeader />
      <PersonalDataForm />
      <OrderStory orders={orders} />
      <LogoutButton />
      <Footer />
    </>
  );
}