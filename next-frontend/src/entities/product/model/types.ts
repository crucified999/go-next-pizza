export type Product = {
  id: number;
  title: string;
  description: string;
  price: number;
  image: string;
  category: string;
  variants?: Variant[];
};

export type Variant = {
  productId: number;
  doughType?: string;
  price: number;
  size: string;
  image: string;
  weight: number;
}