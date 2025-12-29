"use client";

import { Modal } from "@/shared/ui/modal";
import { Pizza, Product, Variant } from "../../model/types";
import { ToppingsList } from "../toppings-list";
import { Button } from "@/shared/ui/button";
import { Variants } from "../variants/variants";
import { ReactNode, useEffect, useMemo, useState } from "react";
import { useToppings } from "@/shared/lib/hooks/useToppings";
import { useRouter } from "next/navigation";
import { useAppDispatch } from "@/app/store";
import { showNotification } from "@/shared/ui/alert/alertSlice";
import { cn } from "@/shared/lib/utils";

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
  onAmountChange?: (amount: string) => void;
  onAdd?: () => void;
  initialSizeIndex?: number;
  selectedToppings?: number[];
  onToggleTopping: (toppingId: number) => void;
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
  onAmountChange,
  onAdd,
  onToggleTopping,
  initialSizeIndex = 0,
  selectedToppings = [],
}: BaseProductModalProps) => {
  const router = useRouter();
  const dispatch = useAppDispatch();
  const [activeSizeIndex, setActiveSizeIndex] = useState(initialSizeIndex);
  const activeSize = variantOptions[activeSizeIndex]?.value || "";

  const handleAddToCart = async () => {
    if (onAdd) {
      await onAdd();
    }

    dispatch(
      showNotification({
        message: `Добавлено:\n
        ${product.title}`,
        type: "success",
      })
    );

    router.back();
  };

  useEffect(() => {
    if (initialSizeIndex >= 0 && initialSizeIndex < variantOptions.length) {
      setActiveSizeIndex(initialSizeIndex);
    }
  }, [initialSizeIndex, variantOptions.length]);

  useEffect(() => {
    if (onAmountChange && variantOptions[activeSizeIndex]?.value) {
      const variantValue = variantOptions[activeSizeIndex].value;

      onAmountChange(variantValue);
    }
  }, [activeSizeIndex, variantOptions, onAmountChange]);

  useEffect(() => {
    if (onSizeChange && activeSize) {
      onSizeChange(activeSize);
    }
  }, [activeSize, onSizeChange]);

  const handleVariantChange = (index: number) => {
    setActiveSizeIndex(index);
  };

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

  return (
    <Modal className="grid grid-cols-[2fr_1.5fr] z-200">
      {renderImage({ activeImage, activeSize })}

      <div className={cn(className, "dark:bg-[#101113]")}>
        <div className="flex-1 overflow-auto">
          <h1 className="text-2xl font-bold">{product.title}</h1>
          {renderContent()}

          {variantOptions.length > 0 && (
            <Variants
              options={variantOptions}
              value={activeSizeIndex}
              onChange={handleVariantChange}
            />
          )}

          {renderAdditionalVariants?.({ activeSize })}

          {product.toppings !== undefined && (
            <ToppingsList
              toppings={product.toppings}
              selectedToppings={selectedToppings}
              onToggleTopping={onToggleTopping}
            />
          )}
        </div>

        <div className="flex items-center justify-center pb-5">
          <Button
            variant="outline"
            className="text-lg h-12 border-none bg-[#FE5F00] text-white hover:bg-[#e55400] w-full"
            onClick={handleAddToCart}
          >
            В корзину за {totalPrice} ₽
          </Button>
        </div>
      </div>
    </Modal>
  );
};
