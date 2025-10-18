"use client";

import { useState } from "react";

import { ArrowRight, ShoppingCart } from "lucide-react";
import { Button } from "@/shared/ui/button";
import { cn } from "@/shared/lib/utils";

export const CartButton = () => {
  const [isHovered, setIsHovered] = useState(false);

  const handleHover = () => {
    setIsHovered(true);
  };

  const handleLeave = () => {
    setIsHovered(false);
  };

  return (
    <Button
      variant="outline"
      onMouseEnter={handleHover}
      onMouseLeave={handleLeave}
      className={cn(
        "p-5 transition-all duration-500 linear",
        isHovered && "bg-[#FE5F00] text-white"
      )}
    >
      {isHovered ? (
        <div className="flex items-center gap-2">
          <span
            className={cn(
              "opacity-0 transition-all duration-500 linear",
              isHovered && "opacity-100"
            )}
          >
            Корзина
          </span>
          <ArrowRight
            width={15}
            className={cn(
              "opacity-0 transition-all duration-500 linear translate-x-[-10px]",
              isHovered && "opacity-100 translate-x-0"
            )}
          />
        </div>
      ) : (
        <ShoppingCart
          width={15}
          className={cn(
            "transition-all duration-500 linear opacity-100",
            isHovered && "opacity-0"
          )}
        />
      )}
    </Button>
  );
};
