import { Combo } from "@/entities/combo/model/types";
import { Product } from "@/entities/product/model/types";

export type OrderStatus = "pending" | "confirmed" | "delivered" | "cancelled";

export type Order = {
  id: number;
  userId: number;
  products?: Product[];
  combos?: Combo[];
  totalPrice: number;
  status: OrderStatus;
  createdAt: Date;
  updatedAt: Date;
};

export type OrderModal = {};
