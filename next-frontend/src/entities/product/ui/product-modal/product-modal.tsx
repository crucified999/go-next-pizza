"use client";

import { Product } from "../../model/types";
import Image from "next/image";
import { useEffect, useMemo, useRef, useState } from "react";
import { BaseProductModal } from "./base-product-modal";
import { useAppDispatch } from "@/app/store";
import { addProduct } from "@/entities/cart/lib/api";
import { addNewProduct } from "@/entities/cart/store/cartSlice";

type ProductModalProps = {
  product: Product;
};

export const ProductModal = ({ product }: ProductModalProps) => {
  const dispatch = useAppDispatch();
  const variants = Array.from(new Set(product.variants?.map((v) => v.size)));
  const initializedRef = useRef(false);

  const [currentAmount, setCurrentAmount] = useState<string | null>(null);
  const [currentSizeIndex, setCurrentSizeIndex] = useState<number | null>(null);

  useEffect(() => {
    if (initializedRef.current) return;

    const savedAmount = localStorage.getItem("amount");
    const savedSizeIndex = localStorage.getItem("product-size-index");

    if (savedAmount) {
      const amount = savedAmount;
      if (amount) {
        setCurrentAmount(amount);
      } else {
        setCurrentAmount(product.amount);
      }
    } else {
      setCurrentAmount(product.amount);
    }

    if (savedSizeIndex) {
      const sizeIndex = Number(savedSizeIndex);
      if (!isNaN(sizeIndex) && sizeIndex >= 0 && sizeIndex < variants.length) {
        setCurrentSizeIndex(sizeIndex);
      } else {
        setCurrentSizeIndex(0);
      }
    } else {
      setCurrentSizeIndex(0);
    }

    initializedRef.current = true;
  }, [product.amount, variants.length]);

  useEffect(() => {
    if (!initializedRef.current || currentAmount === null) return;
    localStorage.setItem("amount", currentAmount);
  }, [currentAmount]);

  useEffect(() => {
    if (!initializedRef.current || currentSizeIndex === null) return;
    localStorage.setItem("product-size-index", currentSizeIndex.toString());
  }, [currentSizeIndex]);

  useEffect(() => {
    return () => {
      localStorage.removeItem("amount");
      localStorage.removeItem("product-size-index");
    };
  }, []);

  const variantOptions = useMemo(() => {
    if (variants.length > 0) {
      return variants.map((v) => ({ value: v, available: true }));
    }

    const amountToShow =
      currentAmount !== null ? currentAmount : product.amount;

    return [
      {
        value:
          product.category !== "drink" && product.category !== "coffee"
            ? amountToShow.toString()
            : amountToShow.toString(),
        available: true,
      },
    ];
  }, [variants, product.category, product.amount, currentAmount]);

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

  const handleAmountChange = (newAmount: string) => {
    setCurrentAmount(newAmount);
  };

  const handleSizeChange = (size: string) => {
    const newIndex = variants.findIndex((v) => v === size);
    if (newIndex !== -1) {
      setCurrentSizeIndex(newIndex);
    }
  };


  const handleAddProduct = async () => {
    const addedProduct = await addProduct({
      productId: product.id,
      amount: currentAmount!,
    });
    
    dispatch(addNewProduct(addedProduct));
  };

  if (currentAmount === null || currentSizeIndex === null) {
    return null;
  }

  return (
    <BaseProductModal
      product={product}
      initialSizeIndex={currentSizeIndex}
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
      onAmountChange={variants.length ? handleAmountChange : undefined}
      onSizeChange={handleSizeChange}
      onAdd={handleAddProduct}
      selectedToppings={[]}
      onToggleTopping={(toppingId: number) => {}}
    />
  );
};


