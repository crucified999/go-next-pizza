"use client";

import { Product } from "@/entities/product/model/types";
import { cn } from "@/shared/lib/utils";
import { Button } from "@/shared/ui/button";
import Image from "next/image";
import { ReactNode } from "react";

type BaseComboCardProps = {
  product: Product;
  renderDetails: () => ReactNode;
  renderActions?: () => ReactNode;
  onReplace: (e?: React.MouseEvent) => void;
  onCardClick?: () => void;
  isActive?: boolean;
  showReplaceButtonWhenActive?: boolean;
};

export const BaseComboCard = ({
  product,
  renderDetails,
  renderActions,
  onReplace,
  onCardClick,
  isActive = false,
  showReplaceButtonWhenActive = false,
}: BaseComboCardProps) => {
  const handleCardClick = () => {
    if (onCardClick) {
      onCardClick();
    }
  };

  return (
    <div
      onClick={handleCardClick}
      className={cn(
        "grid grid-cols-[auto_1fr] gap-5 h-fit bg-white rounded-lg p-2 pb-5 border-1 border-transparent cursor-pointer",
        isActive && "border-1 border-orange-500 rounded-lg"
      )}
    >
      <Image src={product.image} alt={product.title} width={68} height={68} />
      <div>
        <h3 className="text-xl font-bold">{product.title}</h3>
        {renderDetails()}
        <p className="text-sm opacity-80">{product.description}</p>
        {(!isActive || showReplaceButtonWhenActive) && (
          <Button
            variant="outline"
            className="border-none bg-[#FFFAF4] hover:bg-[#ffe7cb] mt-3"
            onClick={(e) => {
              e.stopPropagation();
              onReplace(e);
            }}
          >
            <span>Заменить</span>
          </Button>
        )}
        {renderActions?.()}
      </div>
    </div>
  );
};
