import { Product } from "@/entities/product/model";

export type Combo = {
  id: number;
  title: string;
  description: string;
  price: number;
  image: string;
  defaultProducts: Product[];
  products: Product[];
}