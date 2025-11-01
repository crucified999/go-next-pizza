"use client";

import { Pizza } from "../../model/types";
import { PizzaIngredient } from "../pizza-ingredient/pizza-ingredient";
import { Variants } from "../variants/variants";
import { useEffect, useMemo, useState } from "react";
import { PizzaImage } from "../pizza-image/pizza-image";
import { BaseProductModal } from "../product-modal/base-product-modal";

type PizzaModalProps = {
  pizza: Pizza;
};

const getDoughTypeValue = (doughType: any): number | null => {
  if (typeof doughType === "object" && doughType !== null) {
    if (doughType.Valid && doughType.Int64 !== undefined) {
      return doughType.Int64;
    }
  } else if (typeof doughType === "number") {
    return doughType;
  }
  return null;
};

export const PizzaModal = ({ pizza }: PizzaModalProps) => {
  const sizeVariants = Array.from(new Set(pizza.variants?.map((v) => v.size)));

  const uniqueDoughTypeValues = useMemo(() => {
    const values = new Set<number>();
    (pizza.variants ?? []).forEach((v) => {
      const value = getDoughTypeValue(v.doughType);
      if (value !== null) {
        values.add(value);
      }
    });
    return Array.from(values).sort((a, b) => a - b);
  }, [pizza.variants]);

  const [activeDoughIndex, setActiveDoughIndex] = useState(0);
  const activeDoughValue = uniqueDoughTypeValues[activeDoughIndex] ?? null;

  const [currentSize, setCurrentSize] = useState(sizeVariants[0] || "");

  const doughTypes = useMemo(() => {
    return uniqueDoughTypeValues.map((doughType) =>
      doughType === 1
        ? "Традиционное"
        : doughType === 2
        ? "Тонкое"
        : `Тесто ${doughType}`
    );
  }, [uniqueDoughTypeValues]);

  useEffect(() => {
    const doughOptions = doughTypes.map((d) => {
      const isThinDough = d === "Тонкое";
      const isSmallSize = currentSize === "20 см" || currentSize === "25 см";
      return {
        value: d,
        available: !(isThinDough && isSmallSize),
      };
    });

    const currentOption = doughOptions[activeDoughIndex];
    if (currentOption && !currentOption.available) {
      const firstAvailableIndex = doughOptions.findIndex(
        (opt) => opt.available
      );
      if (firstAvailableIndex !== -1) {
        setActiveDoughIndex(firstAvailableIndex);
      }
    }
  }, [currentSize, doughTypes, activeDoughIndex]);

  const variantOptions = useMemo(() => {
    return sizeVariants.map((v) => ({ value: v, available: true }));
  }, [sizeVariants]);

  const getActiveVariant = (size: string) => {
    return pizza.variants?.find((v) => {
      return (
        v.size === size && getDoughTypeValue(v.doughType) === activeDoughValue
      );
    });
  };

  const getActiveImage = (size: string) => {
    if (!pizza.variants || !size) return pizza.image;

    const match = pizza.variants.find((v) => {
      const variantDoughValue = getDoughTypeValue(v.doughType);
      return (
        v.size === size && variantDoughValue === activeDoughValue && v.image
      );
    });

    return match?.image || pizza.image;
  };

  const getBasePrice = () => 0;

  return (
    <BaseProductModal
      product={pizza}
      renderImage={({ activeImage, activeSize: size }) => {
        const baseSize = 30;
        const parsedSize = parseInt(size, 10) || baseSize;
        const scale = Math.max(0.5, Math.min(1.1, parsedSize / baseSize));
        return <PizzaImage src={activeImage} alt={pizza.title} scale={scale} />;
      }}
      renderContent={() => (
        <div className="text-base list-none">
          {pizza.ingredients.map((i, index) => {
            return index === 0 ? (
              <PizzaIngredient
                key={i.id}
                title={i.title}
                replacable={i.replacable}
                last={index === pizza.ingredients!.length - 1}
              />
            ) : (
              <PizzaIngredient
                key={i.id}
                title={i.title.toLowerCase()}
                replacable={i.replacable}
                last={index === pizza.ingredients!.length - 1}
              />
            );
          })}
        </div>
      )}
      renderAdditionalVariants={({ activeSize }) => {
        const doughOptions = doughTypes.map((d) => {
          const isThinDough = d === "Тонкое";
          const isSmallSize = activeSize === "20 см" || activeSize === "25 см";
          return {
            value: d,
            available: !(isThinDough && isSmallSize),
          };
        });

        return doughOptions.length > 0 ? (
          <Variants
            options={doughOptions}
            value={activeDoughIndex}
            onChange={setActiveDoughIndex}
          />
        ) : null;
      }}
      variantOptions={variantOptions}
      getActiveVariant={getActiveVariant}
      getActiveImage={getActiveImage}
      getBasePrice={getBasePrice}
      onSizeChange={setCurrentSize}
    />
  );
};
