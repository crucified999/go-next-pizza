import { Topping } from "@/entities/product/model/types";
import { useState } from "react";

export const useToppings = () => {
  const [selectedToppings, setSelectedToppings] = useState<Topping[]>([]);

  const toggleTopping = (topping: Topping) => {
    if (selectedToppings.includes(topping)) {
      setSelectedToppings(selectedToppings.filter((t) => t.id !== topping.id));
    } else {
      setSelectedToppings([...selectedToppings, topping]);
    }
  }

  return {
    selectedToppings,
    toggleTopping,
  }
}