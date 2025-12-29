import { configureStore } from "@reduxjs/toolkit";
import { TypedUseSelectorHook, useDispatch, useSelector } from "react-redux";
import { productReducer } from "@/entities/product/store";
import { categoryReducer } from "@/entities/category/store";
import { comboReducer } from "@/entities/combo/store";
import { userReducer } from "@/entities/user/store";
import { cartReducer } from "@/entities/cart/store";
import { notificationReducer } from "@/shared/ui/alert";
import { orderReducer } from "@/entities/order/store";
import appReducer from "./appSlice";

export const store = configureStore({
  reducer: {
    products: productReducer,
    categories: categoryReducer,
    combos: comboReducer,
    user: userReducer,
    app: appReducer,
    cart: cartReducer,
    order: orderReducer,
    notification: notificationReducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export const useAppDispatch = () => useDispatch<AppDispatch>();
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;

export default store;
