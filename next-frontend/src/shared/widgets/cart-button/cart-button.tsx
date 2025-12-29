"use client";

import { useState } from "react";

import { ArrowRight } from "lucide-react";
import { Button } from "@/shared/ui/button";
import { cn } from "@/shared/lib/utils";
import { useAppDispatch, useAppSelector } from "@/app/store";
import { setIsCartOpened } from "@/entities/cart/store/cartSlice";
import { Alert } from "@/shared/ui/alert";

export const CartButton = () => {
  const dispatch = useAppDispatch();
  const cart = useAppSelector((state) => state.cart.cart);
  const [isHovered, setIsHovered] = useState(false);
  const { isVisible, message } = useAppSelector((state) => state.notification);

  const handleHover = () => {
    setIsHovered(true);
  };

  const handleLeave = () => {
    setIsHovered(false);
  };

  const handleClick = () => {
    dispatch(setIsCartOpened());
  };

  return (
    <div className="relative flex w-full justify-end">
      <Button
        variant="outline"
        onClick={handleClick}
        onMouseEnter={handleHover}
        onMouseLeave={handleLeave}
        className={cn(
          "py-5 transition-all duration-150 linear self-end bg-[#FE5F00] text-white dark:text-orange-500 hover:bg-[#e55400]"
        )}
      >
        {cart.totalCount !== 0 ? (
          <div className="flex justify-between items-center gap-2">
            <span>Корзина</span>
            <span className="text-xl text-white opacity-20">|</span>
            <span
              className={cn(
                "w-4 flex justify-center transition-all duration-300 linear translate-x-0",
                isHovered && "translate-x-[2px]"
              )}
            >
              {isHovered ? (
                <ArrowRight width={15} />
              ) : (
                <span>{cart.totalCount}</span>
              )}
            </span>
          </div>
        ) : (
          <span>Корзина</span>
        )}
      </Button>
      {isVisible && <Alert message={message} isVisible={isVisible} />}
    </div>
  );
};
