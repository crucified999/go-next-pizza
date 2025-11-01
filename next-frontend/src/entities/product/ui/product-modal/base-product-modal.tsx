"use client";

import { Modal } from "@/shared/ui/modal";
import { Product, Variant } from "../../model/types";
import { ToppingsList } from "../toppings-list";
import { Button } from "@/shared/ui/button";
import { Variants } from "../variants/variants";
import { ReactNode, useEffect, useMemo, useState } from "react";

type VariantOption = {
  value: string;
  available: boolean;
};

type BaseProductModalProps = {
  product: Product;
  renderImage: (args: { activeImage: string; activeSize: string }) => ReactNode;
  renderContent: () => ReactNode;
  renderAdditionalVariants?: (args: { activeSize: string }) => ReactNode;
  variantOptions: VariantOption[];
  getActiveVariant: (size: string) => Variant | undefined;
  getActiveImage: (size: string) => string;
  getBasePrice: () => number;
  className?: string;
  onSizeChange?: (size: string) => void;
};

export const BaseProductModal = ({
  product,
  renderImage,
  renderContent,
  renderAdditionalVariants,
  variantOptions,
  getActiveVariant,
  getActiveImage,
  getBasePrice,
  className = "grid grid-rows-[1fr_auto] flex-col align-center bg-gray-50 pt-5 px-5 rounded-r-xl gap-5",
  onSizeChange,
}: BaseProductModalProps) => {
  const [activeSizeIndex, setActiveSizeIndex] = useState(0);
  const activeSize = variantOptions[activeSizeIndex]?.value || "";

  useEffect(() => {
    if (onSizeChange && activeSize) {
      onSizeChange(activeSize);
    }
  }, [activeSize, onSizeChange]);

  const [selectedToppings, setSelectedToppings] = useState<number[]>([]);

  const activeVariant = useMemo(() => {
    return getActiveVariant(activeSize);
  }, [activeSize, getActiveVariant]);

  const activeImage = useMemo(() => {
    return getActiveImage(activeSize);
  }, [activeSize, getActiveImage]);

  const totalPrice = useMemo(() => {
    const basePrice = activeVariant?.price || getBasePrice();
    const toppingsPrice = selectedToppings.reduce((sum, toppingId) => {
      const topping = product.toppings?.find((t) => t.id === toppingId);
      return sum + (topping?.price || 0);
    }, 0);
    return basePrice + toppingsPrice;
  }, [activeVariant?.price, getBasePrice, selectedToppings, product.toppings]);

  const toggleTopping = (toppingId: number) => {
    setSelectedToppings((prev) =>
      prev.includes(toppingId)
        ? prev.filter((id) => id !== toppingId)
        : [...prev, toppingId]
    );
  };

  return (
    <Modal className="grid grid-cols-[2fr_1.5fr] z-100">
      {renderImage({ activeImage, activeSize })}

      <div className={className}>
        <div className="flex-1 overflow-auto">
          <h1 className="text-2xl font-bold">{product.title}</h1>
          {renderContent()}

          {variantOptions.length > 0 && (
            <Variants
              options={variantOptions}
              value={activeSizeIndex}
              onChange={setActiveSizeIndex}
            />
          )}

          {renderAdditionalVariants?.({ activeSize })}

          {product.toppings !== undefined && (
            <ToppingsList
              toppings={product.toppings}
              selectedToppings={selectedToppings}
              onToggleTopping={toggleTopping}
            />
          )}
        </div>

        <div className="flex items-center justify-center pb-5">
          <Button
            variant="outline"
            className="text-lg h-12 border-none bg-[#FE5F00] text-white hover:bg-[#e55400] w-full"
          >
            В корзину за {totalPrice} ₽
          </Button>
        </div>
      </div>
    </Modal>
  );
};

