import { createSlice, PayloadAction } from "@reduxjs/toolkit";

interface NotificationState {
  isVisible: boolean;
  message: string;
  type: "success" | "error" | "info";
}

const initialState: NotificationState = {
  isVisible: false,
  message: "",
  type: "success",
};

const notificationSlice = createSlice({
  name: "notification",
  initialState,
  reducers: {
    showNotification: (
      state,
      action: PayloadAction<{ message: string; type?: "success" | "error" | "info" }>
    ) => {
      state.isVisible = true;
      state.message = action.payload.message;
      state.type = action.payload.type || "success";
    },
    hideNotification: (state) => {
      state.isVisible = false;
    },
  },
});

export const { showNotification, hideNotification } = notificationSlice.actions;
export default notificationSlice.reducer;