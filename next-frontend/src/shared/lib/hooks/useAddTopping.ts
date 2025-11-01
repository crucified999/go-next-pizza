import { Topping } from "@/entities/product/model/types";
import { useState } from "react";

export const useAddTopping = () => {
  const [toppings, setToppings] = useState<Topping[]>([]);
  const [isAdded, setIsAdded] = useState(false);

  const addTopping = (topping: Topping) => {
    setIsAdded(!isAdded);
    if (isAdded) {
      setToppings(toppings.filter((t) => t.id !== topping.id));
    } else {
      setToppings([...toppings, topping]);
    }
  }

  return {
    isAdded,
    addTopping,
  }
}