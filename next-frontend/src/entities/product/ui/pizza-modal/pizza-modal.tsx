"use client";

import { Pizza } from "../../model/types";
import { PizzaIngredient } from "../pizza-ingredient/pizza-ingredient";
import { Variants } from "../variants/variants";
import { useEffect, useMemo, useRef, useState } from "react";
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
  const initializedRef = useRef(false);

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

  const [activeDoughIndex, setActiveDoughIndex] = useState<number | null>(null);
  const [currentSize, setCurrentSize] = useState<string | null>(null);

  const currentSizeIndex = useMemo(() => {
    if (currentSize === null) return 0;
    const index = sizeVariants.findIndex(size => size === currentSize);
    return index !== -1 ? index : 0;
  }, [currentSize, sizeVariants]);

  const activeDoughValue = useMemo(() => {
    if (activeDoughIndex === null) return null;
    return uniqueDoughTypeValues[activeDoughIndex] ?? null;
  }, [activeDoughIndex, uniqueDoughTypeValues]);

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
    if (initializedRef.current) return;

    const savedDough = localStorage.getItem('dough');
    const savedSize = localStorage.getItem('size');

    if (savedSize && sizeVariants.includes(savedSize)) {
      setCurrentSize(savedSize);
    } else if (sizeVariants.length > 0) {
      setCurrentSize(sizeVariants[0]);
    }

    if (savedDough) {
      const doughIndex = Number(savedDough);
      if (!isNaN(doughIndex) && doughIndex >= 0 && doughIndex < uniqueDoughTypeValues.length) {
        setActiveDoughIndex(doughIndex);
      } else {
        setActiveDoughIndex(0);
      }
    } else {
      setActiveDoughIndex(0);
    }

    initializedRef.current = true;
  }, [sizeVariants, uniqueDoughTypeValues.length]);

  useEffect(() => {
    if (!initializedRef.current || activeDoughIndex === null) return;
    console.log('PizzaModal: Saving dough to localStorage:', activeDoughIndex);
    localStorage.setItem('dough', activeDoughIndex.toString());
  }, [activeDoughIndex]);

  useEffect(() => {
    if (!initializedRef.current || !currentSize) return;
    console.log('PizzaModal: Saving size to localStorage:', currentSize);
    localStorage.setItem('size', currentSize);
  }, [currentSize]);

  useEffect(() => {
    return () => {
      console.log('PizzaModal: Cleaning up localStorage');
      localStorage.removeItem('size');
      localStorage.removeItem('dough');
    };
  }, []);

  useEffect(() => {
    if (!initializedRef.current || !currentSize || activeDoughIndex === null) return;
    
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
        console.log('PizzaModal: Switching dough due to size constraint');
        setActiveDoughIndex(firstAvailableIndex);
      }
    }
  }, [currentSize, doughTypes, activeDoughIndex]);

  const variantOptions = useMemo(() => {
    return sizeVariants.map((v) => ({ value: v, available: true }));
  }, [sizeVariants]);

  const getActiveVariant = (size: string) => {
    if (activeDoughValue === null) return undefined;
    
    return pizza.variants?.find((v) => {
      return (
        v.size === size && getDoughTypeValue(v.doughType) === activeDoughValue
      );
    });
  };

  const getActiveImage = (size: string) => {
    if (!pizza.variants || !size || activeDoughValue === null) return pizza.image;

    const match = pizza.variants.find((v) => {
      const variantDoughValue = getDoughTypeValue(v.doughType);
      return (
        v.size === size && variantDoughValue === activeDoughValue && v.image
      );
    });

    return match?.image || pizza.image;
  };

  const getBasePrice = () => 0;

  // Не рендерим модалку пока не инициализировались оба состояния
  if (currentSize === null || activeDoughIndex === null) {
    console.log('PizzaModal: Waiting for initialization...', { currentSize, activeDoughIndex });
    return null;
  }

  const handleDoughChange = (index: number) => {
    console.log('PizzaModal: Dough changed to', index);
    setActiveDoughIndex(index);
  };

  return (
    <BaseProductModal
      product={pizza}
      initialSizeIndex={currentSizeIndex}
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
            onChange={handleDoughChange}
          />
        ) : null;
      }}
      variantOptions={variantOptions}
      getActiveVariant={getActiveVariant}
      getActiveImage={getActiveImage}
      getBasePrice={getBasePrice}
      onSizeChange={(size) => {
        console.log('PizzaModal: Size changed to', size);
        setCurrentSize(size);
      }}
    />
  );
};