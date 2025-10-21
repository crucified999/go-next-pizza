'use client';

import { Header } from "@/shared/ui/header";
import { ProductCategoryList } from "@/entities/product/ui/product-category-list/product-category-list";
import { ComboList } from "@/entities/combo/ui/combo-list/combo-list";
import { Footer } from "@/shared/ui/footer/footer";
import { PostHeader } from "@/shared/ui/post-header/post-header";
import { RootState, useAppDispatch, useAppSelector } from "@/app/store";
import { fetchProducts } from "@/entities/product/store/productSlice";
import { useEffect } from "react";
import { fetchCombos } from "@/entities/combo/store/comboSlice";

export const HomePage = () => {
  const dispatch = useAppDispatch();
  const products = useAppSelector(
    (state: RootState) => state.products.products
  );
  const combos = useAppSelector((state) => state.combos.combos);

  useEffect(() => {
    dispatch(fetchProducts());
    dispatch(fetchCombos());
  }, [dispatch]);

  return (
    <>
      <Header />
      <PostHeader />
      <ProductCategoryList category="pizza" products={products}  />
      <ComboList combos={combos} />
      <ProductCategoryList category="snack" products={products} />
      <ProductCategoryList category="shake" products={products} />
      <ProductCategoryList category="coffee" products={products} />
      <ProductCategoryList category="drink" products={products} />
      <ProductCategoryList category="dessert" products={products} />
      <ProductCategoryList category="sauce" products={products} />
      <Footer />
    </>
  );
};
