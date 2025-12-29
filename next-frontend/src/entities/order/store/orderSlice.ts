import { createSlice, createAsyncThunk, PayloadAction } from "@reduxjs/toolkit";
import { getOrders } from "../lib/api";
import type { Order } from "../model/types";

export const fetchOrders = createAsyncThunk("getOrders", getOrders);

type OrderSlice = {
  orders: Order[];
  loading: boolean;
  error: string;
}

const initialState: OrderSlice = {
  orders: [],
  loading: false,
  error: '',
};

export const orderSlice = createSlice({
  name: "order",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchOrders.pending, (state) => {
        state.loading = true;
      })
      .addCase(fetchOrders.fulfilled, (state, action) => {
        state.loading = false;
        state.orders = action.payload || []
      })
      .addCase(fetchOrders.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error as string;
      });
  },
});

export default orderSlice.reducer;
