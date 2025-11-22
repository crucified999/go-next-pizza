"use client";

import { Button } from "@/shared/ui/button";
import React from "react";
import { Variant } from "@/entities/product/model/types";
import Link from "next/link";
import { setIsModalOpened } from "@/app/appSlice";
import { useAppDispatch, useAppSelector } from "@/app/store";

type ProductCardProps = {
  id: number;
  title: string;
  description: string;
  price: number;
  image: string;
  variants: Variant[];
};
export const ProductCard: React.FC<ProductCardProps> = ({
  id,
  title,
  description,
  price,
  image,
  variants,
}) => {
  const isAuth = useAppSelector((state) => state.user.isAuth);
  const href = isAuth ? `/product/${id}` : '/auth';
  const dispatch = useAppDispatch();

  const handleOpenModal = () => {
    dispatch(setIsModalOpened(true));
  };

  return (
    <article className="flex flex-col gap-2 shrink-1 max-w-[300px] h-full">
      <Link href={href} onClick={handleOpenModal}>
        <img
          src={image}
          alt={title}
          className="cursor-pointer max-w-[300px] transition-all duration-300 linear translate-y-0 hover:translate-y-[5px]"
        />
      </Link>
      <h3 className="text-xl font-bold">{title}</h3>
      <p className="text-sm opacity-70 flex-grow">{description}</p>

      <div className="mt-auto flex justify-between items-center">
        <span className="text-lg font-[700]">
          {variants.length > 0 ? `от ${price}` : price} ₽
        </span>
        <Link href={href} onClick={handleOpenModal}>
          <Button
            variant="outline"
            className="border-none bg-[#FFFAF4] hover:bg-[#ffe7cb]"
          >
            {variants.length > 0 ? (
              <span>Выбрать</span>
            ) : (
              <span>В корзину</span>
            )}
          </Button>
        </Link>
      </div>
    </article>
  );
};
