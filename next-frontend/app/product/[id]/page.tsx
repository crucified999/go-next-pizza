import { ProductModal } from "@/entities/product/ui/product-modal";
import { getProductById } from "@/entities/product/lib/api";
import { HomePage } from "@/pages/home-page";
import { PizzaModal } from "@/entities/product/ui/pizza-modal";
import { Pizza } from "@/entities/product/model/types";

interface ProductPageProps {
  params: {
    id: string;
  };
}

export default async function ProductPage({ params }: ProductPageProps) {
  const { id } = params;
  const product = await getProductById(Number(id));

  if (!product) {
    return <div>Продукт не найден</div>;
  }

  return (
    <>
      <HomePage />
      {product.category === "pizza" ? (
        <PizzaModal pizza={product as Pizza} />
      ) : (
        <ProductModal product={product} />
      )}
    </>
  );
}
