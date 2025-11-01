import { configureStore } from "@reduxjs/toolkit";
import { TypedUseSelectorHook, useDispatch, useSelector } from "react-redux";
import { productReducer } from "@/entities/product/store";
import { categoryReducer } from "@/entities/category/store";
import { comboReducer } from "@/entities/combo/store";
import appReducer from "./appSlice";

export const store = configureStore({
  reducer: {
    products: productReducer,
    categories: categoryReducer,
    combos: comboReducer,
    app: appReducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export const useAppDispatch = () => useDispatch<AppDispatch>();
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;

export default store;
