"use client";

import { Product } from "../../model/types";
import Image from "next/image";
import { useMemo } from "react";
import { BaseProductModal } from "./base-product-modal";

type ProductModalProps = {
  product: Product;
};

export const ProductModal = ({ product }: ProductModalProps) => {
  const variants = Array.from(new Set(product.variants?.map((v) => v.size)));

  const variantOptions = useMemo(() => {
    if (variants.length > 0) {
      return variants.map((v) => ({ value: v, available: true }));
    }
    return [
      {
        value:
          product.category !== "drink"
            ? product.amount.toString() + " шт"
            : product.amount.toString() + " л",
        available: true,
      },
    ];
  }, [variants, product.category, product.amount]);

  const getActiveVariant = (size: string) => {
    if (!product.variants || !size) return undefined;
    return product.variants.find((v) => v.size === size);
  };

  const getActiveImage = (size: string) => {
    if (!product.variants || !size) return product.image;
    const match = product.variants.find((v) => v.size === size && v.image);
    return match?.image || product.image;
  };

  const getBasePrice = () => product.price || 0;

  return (
    <BaseProductModal
      product={product}
      renderImage={({ activeImage }) => (
        <div className="flex justify-center items-center px-10">
          <Image
            src={activeImage}
            alt={product.title}
            width={500}
            height={500}
          />
        </div>
      )}
      renderContent={() => (
        <div className="text-base list-none">{product.description}</div>
      )}
      variantOptions={variantOptions}
      getActiveVariant={getActiveVariant}
      getActiveImage={getActiveImage}
      getBasePrice={getBasePrice}
      className="grid grid-rows-[1fr_auto] flex-col align-center bg-gray-50 p-5 rounded-r-xl"
    />
  );
};
