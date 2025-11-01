import { Topping } from "@/entities/product/model/types";
import { useState } from "react";

export const useToppings = (initialToppings: Topping[]) => {
  const [selectedToppings, setSelectedToppings] = useState<number[]>(initialToppings.map((t) => t.id));

  const toggleTopping = (toppingId: number) => {
    setSelectedToppings((prev) =>
      prev.includes(toppingId)
        ? prev.filter((id) => id !== toppingId)
        : [...prev, toppingId]
    );
  };

  return {
    selectedToppings,
    toggleTopping,
  }
}