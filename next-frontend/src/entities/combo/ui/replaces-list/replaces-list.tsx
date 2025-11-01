import { Product } from "@/entities/product/model/types";
import { ReplaceCard } from "../replace-card";

type ReplacesListProps = {
  products: Product[];
};

export const ReplacesList = ({ products }: ReplacesListProps) => {
  return (
    <ul className="grid grid-cols-3 p-5 gap-10 h-full">
      {products.map((product) => (
        <ReplaceCard key={product.id} product={product} />
      ))}
    </ul>
  );
};
