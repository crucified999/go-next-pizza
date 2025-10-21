import { Category } from "../model";

export const getCategories = (): Promise<Category[]> => {
  return fetch(`${process.env.NEXT_PUBLIC_API_URL}/categories`)
    .then((res) => res.json())
    .then((data) => data.map((category: Category) => ({
      id: category.id,
      title: category.title,
    })));
};