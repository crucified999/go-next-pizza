import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import { getProductById, getProducts } from "../lib/api";
import { Product } from "../model/types";

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

type ProductState = {
  products: Product[];
  loading: boolean;
  error: string | null;
}

const initialState: ProductState = {
  products: [],
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
    });
  },
});

export default productSlice.reducer;