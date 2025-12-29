import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  isCartOpened: false,
};


export const appSlice = createSlice({
  name: "app",
  initialState,
  reducers: {
    setIsCartOpened: (state) => {
      state.isCartOpened = !state.isCartOpened;
    },
  },
});

export const { setIsCartOpened } = appSlice.actions;

export default appSlice.reducer;
