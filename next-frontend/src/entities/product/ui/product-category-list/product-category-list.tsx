"use client";

import React, { useEffect } from "react";
import { ProductCard } from "@/entities/product/ui/product-card/product-card";
import { RootState, useAppDispatch, useAppSelector } from "@/app/store";
import { fetchProducts } from "@/entities/product/store/productSlice";
import { getCategoryName } from "@/entities/product/lib/constants";

type ProductCategoryListProps = {
  category: string;
};

export const ProductCategoryList: React.FC<ProductCategoryListProps> = ({
  category,
}) => {
  const products = useAppSelector(
    (state: RootState) => state.products.products
  ).filter((product) => product.category === category);
  const dispatch = useAppDispatch();

  console.log(products);

  useEffect(() => {
    dispatch(fetchProducts());
  }, [dispatch]);

  return (
    <div className="mt-8 pt-5">
      <h2 className="font-[800] text-4xl leading-[100%]">
        {getCategoryName(category)}
      </h2>
      <ul className="grid grid-cols-4 gap-25 py-20">
        {products.map((product) => (
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
