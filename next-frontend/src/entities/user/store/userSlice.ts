import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import { sendCode, verifyCode, checkAuth } from "../lib/api";
import { Order } from "@/entities/order/model/types";

export const fetchAuth = createAsyncThunk(
  "user/auth",
  async (phone: string) => {
    return await sendCode(phone);
  }
);

export const verifyAuth = createAsyncThunk(
  "user/verify",
  async (code: string) => {
    return await verifyCode(code);
  }
);

export const checkUserAuth = createAsyncThunk("user/checkAuth", async () => {
  return await checkAuth();
});

type UserState = {
  id: number;
  isAuth: boolean;
  phone: string;
  name: string;
  email: string;
  orders: Order[];
  isLoading: boolean;
};

const initialState: UserState = {
  id: 0,
  isAuth: false,
  phone: "",
  name: "",
  email: "",
  orders: [],
  isLoading: false,
};

const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {
    setAuth: (state, action) => {
      state.isAuth = action.payload;
    },
    logout: (state) => {
      state.isAuth = false;
      state.orders = [];
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchAuth.pending, (state) => {
        state.isLoading = true;
      })
      .addCase(fetchAuth.fulfilled, (state) => {
        state.isLoading = false;
      })
      .addCase(fetchAuth.rejected, (state) => {
        state.isAuth = false;
        state.isLoading = false;
      })
      .addCase(verifyAuth.pending, (state) => {
        state.isLoading = true;
      })
      .addCase(verifyAuth.fulfilled, (state, action) => {
        if (action.payload.verified) {
          state.isAuth = true;
        }
        state.isLoading = false;
      })
      .addCase(verifyAuth.rejected, (state) => {
        state.isAuth = false;
        state.isLoading = false;
      })
      .addCase(checkUserAuth.fulfilled, (state, action) => {
        state.isAuth = action.payload.authenticated;
        state.id = action.payload.id
        state.phone = action.payload.phone;
        state.name = action.payload.name;
        state.email = action.payload.email;
        state.isLoading = false;
        console.log(state.id, state.name, state.email);
      })
      .addCase(checkUserAuth.rejected, (state) => {
        state.isAuth = false;
        state.isLoading = false;
      });
  },
});

export const { setAuth, logout } = userSlice.actions;
export default userSlice.reducer;
