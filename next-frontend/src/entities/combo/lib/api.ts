import { Combo } from "../model";

const BASE_URL = process.env.NEXT_PUBLIC_API_URL;

export const getCombos = (): Promise<Combo[]> => {
  return fetch(`${BASE_URL}/combos`)
    .then((res) => res.json())
    // .then((data) => data.map((combo: Combo) => ({
    //   id: combo.id,
    //   title: combo.title,
    //   description: combo.description,
    //   price: combo.price,
    //   image: combo.image,
    //   products: combo.products,
    // })));
};

export const getComboById = (id: number): Promise<Combo> => {
  return fetch(`${BASE_URL}/combos/${id}`)
    .then((res) => res.json())
    // .then((combo: Combo) => ({
    //   id: combo.id,
    //   title: combo.title,
    //   description: combo.description,
    //   price: combo.price,
    //   image: combo.image,
    //   products: combo.products,
    // }));
};