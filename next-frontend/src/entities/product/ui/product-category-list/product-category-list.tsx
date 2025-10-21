"use client";

import React, { useEffect } from "react";
import { ProductCard } from "@/entities/product/ui/product-card/product-card";
import { getCategoryName } from "@/entities/product/lib/constants";
import { useInView } from "react-intersection-observer";
import {
  setCurrentCategoryAutomatically,
  resetManualSelection,
} from "@/entities/category/store/categorySlice";
import { Product } from "@/entities/product/model/index";
import { useAppDispatch } from "@/app/store";

type ProductCategoryListProps = {
  category: string;
  products: Product[];
};

export const ProductCategoryList: React.FC<ProductCategoryListProps> = ({
  category,
  products,
}) => {
  const dispatch = useAppDispatch();
  const { ref, inView } = useInView({
    threshold: 0.1,
  });

  useEffect(() => {
    if (inView) {
      dispatch(setCurrentCategoryAutomatically(getCategoryName(category)));
    }
  }, [inView, category, dispatch]);

  useEffect(() => {
    if (inView) {
      const timer = setTimeout(() => {
        dispatch(resetManualSelection());
      }, 1000);
      return () => clearTimeout(timer);
    }
  }, [inView, dispatch]);

  return (
    <div
      className="mt-8 pt-5 scroll-mt-15"
      id={getCategoryName(category)}
      ref={ref}
    >
      <h2 className="font-[800] text-4xl leading-[100%]">
        {getCategoryName(category)}
      </h2>
      <ul className="grid grid-cols-4 gap-25 py-20">
        {products
          .filter((product) => product.category === category)
          .map((product) => (
            <li key={product.id}>
              <ProductCard
                id={product.id}
                title={product.title}
                description={product.description}
                price={product.price}
                image={product.image}
                variants={product.variants || []}
              />
            </li>
          ))}
      </ul>
    </div>
  );
};
