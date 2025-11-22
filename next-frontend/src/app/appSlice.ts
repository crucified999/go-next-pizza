import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  isModalOpened: false,
};


export const appSlice = createSlice({
  name: "app",
  initialState,
  reducers: {
    setIsModalOpened: (state, action) => {
      state.isModalOpened = action.payload;
    },
  },
});

export const { setIsModalOpened } = appSlice.actions;

export default appSlice.reducer;
