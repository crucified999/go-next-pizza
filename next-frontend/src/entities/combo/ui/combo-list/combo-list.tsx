"use client";

import { useAppDispatch } from "@/app/store";
import { useEffect, useState } from "react";
import { ProductCard } from "@/entities/product/ui/product-card";
import { Combo } from "../../model/combo";
import { useInView } from "react-intersection-observer";
import { setCurrentCategoryAutomatically } from "@/entities/category/store/categorySlice";

type ComboListProps = {
  combos: Combo[];
};

export const ComboList = ({ combos }: ComboListProps) => {
  const [isInitialized, setIsInitialized] = useState(false);
  const { ref, inView } = useInView({
    threshold: 0.3,
  });
  const dispatch = useAppDispatch();

  useEffect(() => {
    const timer = setTimeout(() => {
      setIsInitialized(true);
    }, 1000);
    return () => clearTimeout(timer);
  }, []);

  useEffect(() => {
    const timer = setTimeout(
      () => {
        if (isInitialized && inView) {
          dispatch(setCurrentCategoryAutomatically("Комбо"));
          localStorage.setItem('category', 'Комбо');
        }
      },
      200
    );

    return () => clearTimeout(timer);
    
  }, [inView, dispatch, isInitialized]);

  return (
    <div className="mt-8 pt-5 scroll-mt-15" id="Комбо" ref={ref}>
      <h2 className="font-[800] text-4xl leading-[100%]">Комбо</h2>
      <ul className="grid grid-cols-4 gap-25 py-20">
        {combos.map((combo) => (
          <li key={combo.id}>
            <ProductCard
              id={combo.id}
              title={combo.title}
              description={combo.description}
              price={combo.price}
              image={combo.image}
              variants={combo.products
                .map((product) => product.variants)
                .flat()
                .filter((variant) => variant !== undefined)}
            />
          </li>
        ))}
      </ul>
    </div>
  );
};
