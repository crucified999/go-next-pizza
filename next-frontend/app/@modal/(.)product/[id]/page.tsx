import { ProductModal } from "@/entities/product/ui/product-modal";
import { getProductById } from "@/entities/product/lib/api";
import { PizzaModal } from "@/entities/product/ui/pizza-modal";
import { Pizza } from "@/entities/product/model/types";

export default async function ProductModalPage({
  params,
}: {
  params: { id: string };
}) {
  const { id } = params;
  const product = await getProductById(Number(id));

  if (!product) {
    return <div>Продукт не найден</div>;
  }

  return product.category === "pizza" ? (
    <PizzaModal pizza={product as Pizza} />
  ) : (
    <ProductModal product={product} />
  );
}
