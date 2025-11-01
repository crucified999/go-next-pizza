import { Pizza } from "@/entities/product/model/types";
import { Button } from "@/shared/ui/button";
import { BaseComboCard } from "../combo-product-card/base-combo-card";

type ComboPizzaCardProps = {
  product: Pizza;
  onReplace: () => void;
  onChoose: () => void;
  onCardClick?: () => void;
  isActive?: boolean;
  isInChangeMode?: boolean;
};

const getStringValue = (value: any): string => {
  if (typeof value === "object" && value !== null) {
    if (value.Valid && value.String !== undefined) {
      return value.String;
    }
  } else if (typeof value === "string") {
    return value;
  }
  return "";
};

const getDoughTypeValue = (doughType: any): number | null => {
  if (typeof doughType === "object" && doughType !== null) {
    if (doughType.Valid && doughType.Int64 !== undefined) {
      return doughType.Int64;
    }
  } else if (typeof doughType === "number") {
    return doughType;
  }
  return null;
};

export const ComboPizzaCard = ({
  product,
  onReplace,
  onChoose,
  onCardClick,
  isActive,
  isInChangeMode,
}: ComboPizzaCardProps) => {
  const variant = product.variants?.[0];
  const size = variant ? getStringValue(variant.size) : "";
  const doughTypeValue = variant ? getDoughTypeValue(variant.doughType) : null;
  const doughType =
    doughTypeValue === 1
      ? "традиционное"
      : doughTypeValue === 2
      ? "тонкое"
      : "";
  const weight = variant?.weight || 0;

  return (
    <BaseComboCard
      product={product}
      isActive={isActive}
      renderDetails={() => (
        <span className="text-sm opacity-80">{`${size}, ${doughType} тесто, ${weight} г`}</span>
      )}
      renderActions={() =>
        !isInChangeMode ? (
          <Button
            variant="outline"
            className="border-none bg-none mt-3 z-1"
            onClick={(e) => {
              e.stopPropagation();
              onChoose();
            }}
          >
            <span>Изменить состав</span>
          </Button>
        ) : null
      }
      onReplace={(e) => {
        e?.stopPropagation();
        onReplace();
      }}
      onCardClick={onCardClick}
    />
  );
};
