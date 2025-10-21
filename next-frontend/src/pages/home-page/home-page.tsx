import { Header } from "@/shared/ui/header";
import { ProductCategoryList } from "@/entities/product/ui/product-category-list/product-category-list";
import { CartButton } from "@/shared/widgets/cart-button";
import { CategoryList } from "@/entities/category/ui/category-list";
import { ProductCard } from "@/entities/product/ui/product-card/product-card";
import { ComboList } from "@/entities/combo/ui/combo-list/combo-list";

export const HomePage = () => {
  return (
    <>
      <Header />
      <div className="flex justify-between mb-10">
        <CategoryList />
        <CartButton />
      </div>
      <ProductCategoryList category="pizza" />
      <ComboList />
      <ProductCategoryList category="snack" />
      <ProductCategoryList category="shakes" />
      <ProductCategoryList category="drink" />
      <ProductCategoryList category="dessert" />
      <ProductCategoryList category="sauce" />
    </>
  );
};
