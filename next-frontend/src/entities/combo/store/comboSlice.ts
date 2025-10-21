import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import { getCombos } from "../lib/api";
import { Combo } from "../model";

interface ComboState {
  combos: Combo[];
  loading: boolean;
  error: string | null;
}

const initialState: ComboState = {
  combos: [],
  loading: false,
  error: null,
};

export const fetchCombos = createAsyncThunk("combo/fetchCombos", getCombos);

const comboSlice = createSlice({
  name: "combo",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(fetchCombos.pending, (state) => {
      state.loading = true;
    });
    builder.addCase(fetchCombos.fulfilled, (state, action) => {
      state.combos = action.payload;
      state.loading = false;
    });
    builder.addCase(fetchCombos.rejected, (state, action) => {
      state.error = action.error.message || null;
      state.loading = false;
    });
  },
});

export default comboSlice.reducer;