import { Combo } from "../model";

export const getCombos = (): Promise<Combo[]> => {
  return fetch(`${process.env.NEXT_PUBLIC_API_URL}/combos`)
    .then((res) => res.json())
    .then((data) => data.map((combo: Combo) => ({
      id: combo.id,
      title: combo.title,
      description: combo.description,
      price: combo.price,
      image: combo.image,
      defaultProducts: combo.defaultProducts,
      products: combo.products,
    })));
};