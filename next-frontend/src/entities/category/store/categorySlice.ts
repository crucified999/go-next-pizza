import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import { Category } from "../model";
import { getCategories } from "../lib/api";

interface CategoryState {
  categories: Category[];
  currentCategory: string;
  loading: boolean;
  error: string | null;
  isManuallySelected: boolean;
}

const CATEGORY = typeof window !== "undefined" && localStorage.getItem("category") || "";

const initialState: CategoryState = {
  categories: [],
  currentCategory: CATEGORY,
  loading: false,
  error: null,
  isManuallySelected: false,
};

export const fetchCategories = createAsyncThunk(
  "category/fetchCategories",
  getCategories
);

const categorySlice = createSlice({
  name: "category",
  initialState,
  reducers: {
    setCurrentCategory: (state, action) => {
      state.currentCategory = action.payload;
    },
    setCurrentCategoryManually: (state, action) => {
      state.currentCategory = action.payload;
      state.isManuallySelected = true;
    },
    setCurrentCategoryAutomatically: (state, action) => {
      if (!state.isManuallySelected) {
        state.currentCategory = action.payload;
      }
    },
    resetManualSelection: (state) => {
      state.isManuallySelected = false;
    },
  },
  extraReducers: (builder) => {
    builder.addCase(fetchCategories.pending, (state) => {
      state.loading = true;
    });
    builder.addCase(fetchCategories.fulfilled, (state, action) => {
      state.categories = action.payload;
      state.loading = false;
    });
    builder.addCase(fetchCategories.rejected, (state, action) => {
      state.error = action.error.message || null;
      state.loading = false;
    });
  },
});

export const {
  setCurrentCategory,
  setCurrentCategoryManually,
  setCurrentCategoryAutomatically,
  resetManualSelection,
} = categorySlice.actions;

export default categorySlice.reducer;
