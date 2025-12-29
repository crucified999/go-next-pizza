"use client";

import { useAppDispatch } from "@/app/store";
import { addProduct, deleteProduct } from "@/entities/cart/lib/api";
import { ProductInCart } from "@/entities/cart/model/types";
import { addNewProduct, deleteCartProduct, deleteProductCompletely } from "@/entities/cart/store/cartSlice";
import { X } from "lucide-react";
import React, { useEffect, useState } from "react";

type CartProductProps = {
  cartItem: ProductInCart;
};

export const CartProduct: React.FC<CartProductProps> = ({ cartItem }) => {
  const dispatch = useAppDispatch();
  const [count, setCount] = useState(cartItem.count);

  const handleAddProduct = async () => {
    const addedProduct = await addProduct({
      productId: cartItem.product.productId,
      amount: cartItem.product.size,
    });

    dispatch(addNewProduct(addedProduct));
    setCount((prev) => prev + 1);
  };

  const handleDeleteProduct = async () => {
    await deleteProduct({
      productId: cartItem.product.productId,
      amount: cartItem.product.size,
      action: "delete",
    });

    dispatch(deleteCartProduct(cartItem.product));
    setCount((prev) => prev - 1);
  };

  const handleDeleteCompletely = async () => {
    await deleteProduct({
      productId: cartItem.product.productId,
      amount: cartItem.product.size,
      action: "delete-completely",
    });

    dispatch(deleteProductCompletely(cartItem));
    setCount(0);
  }

  useEffect(() => {
    setCount(cartItem.count);
  }, [cartItem.count]);

  return count > 0 && (
    <article className="bg-white w-auto p-5 relative rounded-2xl dark:bg-[#101113] dark:border-1 dark:border-white">
      <button onClick={handleDeleteCompletely} className="absolute top-5 right-10 cursor-pointer">
        <X width={20} height={20} />
      </button>
      <div className="grid grid-cols-[auto_1fr] gap-5 mb-3">
        <img
          className="w-16 h-16"
          src={cartItem.product.image}
          alt={cartItem.product.title}
        />
        <div>
          <h4 className="text-xl font-[600]">{cartItem.product.title}</h4>
          <span className="text-sm text-black/50">{cartItem.product.size}</span>
        </div>
      </div>
      <hr />
      <div className="flex justify-between mt-3">
        <span className="font-[600]">{cartItem.product.price} ₽</span>
        <div className="flex gap-2 justify-center items-center">
          <button className="text-orange-500 cursor-pointer hover:text-[#e55400] transition-colors duration-150 linear text-[15px]">
            Изменить
          </button>
          <div className="flex items-center justify-around gap-6 bg-gray-100 dark:bg-black border-none rounded-2xl py-1 px-3">
            <button
              onClick={handleDeleteProduct}
              className="cursor-pointer text-xl"
            >
              –
            </button>
            <span>{count}</span>
            <button
              onClick={handleAddProduct}
              className="cursor-pointer text-xl"
            >
              +
            </button>
          </div>
        </div>
      </div>
    </article>
  );
};
