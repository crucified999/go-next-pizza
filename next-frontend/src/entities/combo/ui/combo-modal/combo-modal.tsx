"use client";

import { Modal } from "@/shared/ui/modal";
import { Combo } from "../../model";
import Image from "next/image";
import { ComboProductCard } from "../combo-product-card";
import { ComboPizzaCard } from "../combo-pizza-card";
import { Pizza } from "@/entities/product/model/types";
import { ReplacesList } from "../replaces-list";
import { useContent } from "@/shared/lib/hooks/useContent";
import { ComboChoosePizza } from "../combo-choose-pizza";
import { useState } from "react";
import { cn } from "@/shared/lib/utils";

type ComboModalProps = {
  combo: Combo;
};

type ContentMode = "default" | "replace" | "choose";

export const ComboModal = ({ combo }: ComboModalProps) => {
  const defaultImage = (
    <Image src={combo.image} alt={combo.title} width={500} height={500} />
  );
  const [content, setContent, resetContent] = useContent(defaultImage);
  const [activeProductId, setActiveProductId] = useState<number | null>(null);
  const [contentMode, setContentMode] = useState<ContentMode>("default");

  const handleChoose = (productId: number, pizza: Pizza, category: string) => {
    setActiveProductId(productId);
    setContentMode("choose");
    setContent(
      <ComboChoosePizza
        pizza={pizza}
        onReplace={() => handleReplace(productId, category)}
      />
    );
  };

  const handleReplace = (productId: number, category: string) => {
    setActiveProductId(productId);
    setContentMode("replace");
    setContent(
      <ReplacesList
        products={combo.products.filter((p) => p.category === category)}
      />
    );
  };

  const handleCardClick = (productId: number, category: string) => {
    if (activeProductId === productId) {
      setActiveProductId(null);
      setContentMode("default");
      resetContent();
    } else {
      handleReplace(productId, category);
    }
  };

  return (
    <Modal className="grid grid-cols-[2fr_1.5fr]">
      <div
        className={cn(
          "flex justify-center py-5",
          contentMode === "default" && "items-center"
        )}
      >
        {content}
      </div>
      <div className="grid grid-rows-[1fr_auto] flex-col gap-2 bg-gray-100 rounded-r-xl p-5 overflow-auto">
        <div className="flex flex-col gap-5 w-full">
          <div className="flex flex-col">
            <h1 className="text-2xl font-bold">{combo.title}</h1>
            <p className="text-sm opacity-80">{combo.description}</p>
          </div>

          {combo.defaultProducts.map((product) =>
            product.category === "pizza" ? (
              <ComboPizzaCard
                key={product.id}
                product={product as Pizza}
                isActive={activeProductId === product.id}
                isInChangeMode={contentMode === "choose"}
                onReplace={() => handleReplace(product.id, product.category)}
                onChoose={() =>
                  handleChoose(product.id, product as Pizza, product.category)
                }
                onCardClick={() =>
                  handleCardClick(product.id, product.category)
                }
              />
            ) : (
              <ComboProductCard
                key={product.id}
                product={product}
                isActive={activeProductId === product.id}
                onReplace={() => handleReplace(product.id, product.category)}
                onCardClick={() =>
                  handleCardClick(product.id, product.category)
                }
              />
            )
          )}
        </div>
      </div>
    </Modal>
  );
};
