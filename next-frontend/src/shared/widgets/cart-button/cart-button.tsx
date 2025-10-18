"use client";

import { useState } from "react";

import { ArrowRight } from "lucide-react";
import { Button } from "@/shared/ui/button";
import { cn } from "@/shared/lib/utils";

export const CartButton = () => {
  const cartItems = [""]; // TODO: Взять из стора
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
        "py-5 transition-all duration-150 linear self-end bg-[#FE5F00] text-white hover:bg-[#a03500]"
      )}
    >
      {
      cartItems.length !== 0 ?
       <div className="flex justify-between items-center gap-2">
          <span>Корзина</span>
          <span className="text-xl text-white opacity-20">|</span> 
          <span className={cn("w-4 flex justify-center transition-all duration-300 linear translate-x-0", isHovered && "translate-x-[2px]")}>
            {isHovered ? <ArrowRight width={15} /> : <span>{cartItems.length}</span>}
          </span>
       </div>
         : 
        <span>Корзина</span>
        }

        {/* {isHovered && <ArrowRight width={15} className={cn("transition-all self-center -translate-x-2 duration-300", isHovered && "translate-x-0")} />} */}
      
      
    </Button>
  );
};
