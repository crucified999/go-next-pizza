"use client";

import { Product } from "@/entities/product/model/types";
import { cn } from "@/shared/lib/utils";
import Image from "next/image";
import { useState } from "react";

type ReplaceCardProps = {
  product: Product;
};

export const ReplaceCard = ({ product }: ReplaceCardProps) => {
  const [isActive, setIsActive] = useState(false);

  return (
    <li
      className={cn(
        "cursor-pointer p-2 h-full max-h-[200px] flex",
        isActive && "border-1 border-orange-500 rounded-lg"
      )}
      onClick={() => setIsActive(!isActive)}
    >
      <article className="flex flex-col items-center justify-between w-full">
        <Image
          src={product.image}
          alt={product.title}
          width={140}
          height={140}
          className="flex-shrink-0 tranistion-transform duration-150 linear hover:scale-[0.95]"
        />
        <span className="text-center font-bold mt-2 flex-grow flex items-center justify-center">
          {product.title}
        </span>
      </article>
    </li>
  );
};
