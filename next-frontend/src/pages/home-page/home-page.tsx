import { Header } from "@/shared/ui/header";
import { ProductCategoryList } from "@/shared/ui/product-category-list/product-category-list";
import { CartButton } from "@/shared/widgets/cart-button";
import { CategoryList } from "@/shared/widgets/category-list";
import { ProductCard } from "@/shared/widgets/product-card/product-card";

export const HomePage = () => {
  return (
    <>
      <Header />
      <div className="flex justify-between">
        <CategoryList />
        <CartButton />
      </div>
      <ProductCategoryList category="Пиццы" />
    </>
  );
}