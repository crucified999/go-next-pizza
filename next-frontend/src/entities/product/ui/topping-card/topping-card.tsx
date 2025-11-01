"use client";

import { cn } from "@/shared/lib/utils";
import { Topping } from "../../model/types";
import { CircleCheckBig } from "lucide-react";

type ToppingCardProps = {
  topping: Topping;
  isSelected: boolean;
  onToggle: () => void;
};

export const ToppingCard = ({
  topping,
  isSelected,
  onToggle,
}: ToppingCardProps) => {
  return (
    <button
      onClick={onToggle}
      className={cn(
        "relative text-sm cursor-pointer h-[160px] bg-white flex flex-col justify-between items-center gap-2 shadow-[1px_1px_10px_rgba(0,0,0,0.15)] rounded-lg p-2 hover:shadow-sm transition-all duration-150 linear w-full border-1 border-white",
        isSelected && "border-orange-500 border-1"
      )}
    >
      <img
        src={topping.image}
        alt={topping.title}
        className="w-16 h-16 object-contain"
      />
      <p className="px-2 text-center leading-tight overflow-hidden text-ellipsis">
        {topping.title}
      </p>
      <span className="text-lg font-bold">{topping.price} ₽</span>
      {isSelected && (
        <CircleCheckBig
          width={20}
          height={20}
          className="absolute top-1 right-2 text-orange-500"
        />
      )}
    </button>
  );
};
