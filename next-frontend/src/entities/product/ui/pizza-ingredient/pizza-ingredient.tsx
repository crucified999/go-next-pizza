"use client";

import React, { useState } from "react";
import { cn } from "@/shared/lib/utils";
import { CircleXIcon, CornerUpLeft, X } from "lucide-react";

export type PizzaIngredientProps = {
  title: string;
  replacable: boolean;
  last: boolean;
};

export const PizzaIngredient: React.FC<PizzaIngredientProps> = ({
  title,
  replacable,
  last,
}) => {
  const [isActive, setIsActive] = useState<boolean>(true);

  return replacable ? (
    <button
      className={cn("inline-flex mr-1 cursor-pointer")}
      onClick={() => setIsActive(!isActive)}
    >
      <span
        className={cn(
          isActive && replacable && "underline decoration-dotted",
          !isActive && replacable && "line-through"
        )}
      >
        {title}&nbsp;
      </span>
      {isActive && replacable ? (
        <X
          width={14}
          height={14}
          className="border-1 border-black/60 p-0.5 rounded-full"
        />
      ) : (
        <CornerUpLeft
          width={14}
          height={14}
          className="border-1 border-black/60 p-0.5 rounded-full"
        />
      )}
      {last ? "" : ","}
    </button>
  ) : (
    <span className="mr-1">
      {title}
      {last ? "" : ","}&nbsp;
    </span>
  );
};
