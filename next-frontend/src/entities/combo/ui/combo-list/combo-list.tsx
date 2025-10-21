"use client";

import { useAppDispatch, useAppSelector } from "@/app/store";
import { fetchCombos } from "../../store/comboSlice";
import { useEffect } from "react";
import { ProductCard } from "@/entities/product/ui/product-card";

export const ComboList = () => {
  const combos = useAppSelector((state) => state.combos.combos);
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(fetchCombos());
  }, [dispatch]);

  return (
    <div className="mt-8 pt-5">
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
              variants={combo.products.map((product) => product.variants).flat().filter((variant) => variant !== undefined)}
            />
          </li>
        ))}
      </ul>
    </div>
  );
};
