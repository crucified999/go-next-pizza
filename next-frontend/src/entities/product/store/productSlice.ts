import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import {
  createCustomPizza,
  getCustomPizza,
  getProductById,
  getProducts,
  updateCustomPizza,
} from "../lib/api";
import { CustomPizza, Product } from "../model/types";

export const fetchProducts = createAsyncThunk<Product[]>(
  "products/fetchProducts",
  async () => {
    const response = await getProducts();
    return response;
  }
);

export const fetchProductById = createAsyncThunk<Product, number>(
  "products/fetchProductById",
  async (id) => {
    const response = await getProductById(id);
    return response;
  }
);

export const fetchCustomPizza = createAsyncThunk(
  "products/fetchCustomPizza",
  getCustomPizza
);

export const fetchCreateCustomPizza = createAsyncThunk(
  "products/fetchCreateCustomPizza",
  createCustomPizza
);

export const fetchUpdateCustomPizza = createAsyncThunk(
  "products/fetchUpdateCustomPizza",
  updateCustomPizza
);

type ProductState = {
  products: Product[];
  customPizzas: CustomPizza[];
  currentCustomPizza: CustomPizza;
  loading: boolean;
  error: string | null;
};

const initialState: ProductState = {
  products: [],
  customPizzas: [],
  currentCustomPizza: {
    id: 0,
    ingredients: [],
    totalPrice: 0,
    dough: "Традиционное",
    size: "20 см",
  },
  loading: false,
  error: null,
};

const productSlice = createSlice({
  name: "products",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchProducts.pending, (state) => {
        state.loading = true;
      })
      .addCase(fetchProducts.fulfilled, (state, action) => {
        state.products = action.payload;
        state.loading = false;
      })
      .addCase(fetchProducts.rejected, (state, action) => {
        state.error = action.error.message || "Failed to fetch products";
        state.loading = false;
      })
      .addCase(fetchProductById.pending, (state) => {
        state.loading = true;
      })
      .addCase(fetchProductById.fulfilled, (state, action) => {
        state.products = [...state.products, action.payload];
        state.loading = false;
      })
      .addCase(fetchProductById.rejected, (state, action) => {
        state.error = action.error.message || "Failed to fetch product";
        state.loading = false;
      })
      .addCase(fetchCreateCustomPizza.pending, (state) => {
        state.loading = true;
      })
      .addCase(fetchCreateCustomPizza.fulfilled, (state, action) => {
        state.customPizzas = [...state.customPizzas, action.payload];
        state.loading = false;
      })
      .addCase(fetchCreateCustomPizza.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error as string;
      })
      .addCase(fetchUpdateCustomPizza.pending, (state) => {
        state.loading = true;
      })
      .addCase(fetchUpdateCustomPizza.fulfilled, (state, action) => {
        for (let i = 0; i < state.customPizzas.length; i++) {
          if (state.customPizzas[i].id === action.payload.id) {
            state.customPizzas[i] = action.payload;
            break;
          }
        }
        state.loading = false;
      })
      .addCase(fetchUpdateCustomPizza.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error as string;
      })
      .addCase(fetchCustomPizza.pending, (state) => {
        state.loading = true;
      })
      .addCase(fetchCustomPizza.fulfilled, (state, action) => {
        state.currentCustomPizza = action.payload;
        state.loading = false;
      })
      .addCase(fetchCustomPizza.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error as string;
      })
  },
});

export default productSlice.reducer;
