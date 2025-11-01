import { Product } from "@/entities/product/model/types";
import { BaseComboCard } from "./base-combo-card";

type ComboProductCardProps = {
  product: Product;
  onReplace: () => void;
  onCardClick?: () => void;
  isActive?: boolean;
};

export const ComboProductCard = ({
  product,
  onReplace,
  onCardClick,
  isActive,
}: ComboProductCardProps) => {
  return (
    <BaseComboCard
      product={product}
      isActive={isActive}
      renderDetails={() => (
        <span className="text-sm opacity-80">{`${product.amount} шт, ${product.weight} г`}</span>
      )}
      onReplace={(e) => {
        e?.stopPropagation();
        onReplace();
      }}
      onCardClick={onCardClick}
    />
  );
};
