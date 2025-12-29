import {
  Action,
  createAsyncThunk,
  createSlice,
  PayloadAction,
} from "@reduxjs/toolkit";
import { Cart, ProductInCart, PizzaInCart } from "../model/types";
import { getCart, refreshCart } from "../lib/api";
import { PizzaVariant, ProductVariant } from "@/entities/product/model";

type CartSlice = {
  isCartOpened: boolean;
  cart: Cart;
  loading: boolean;
  error: string;
};

export const fetchCart = createAsyncThunk("getCart", async () => {
  const cart = await getCart();

  return {
    ...cart,
    products: cart.products || [],
    pizzas: (cart.pizzas || []).map(pizza => ({
      ...pizza,
      pizza: {
        ...pizza.pizza,
        toppings: pizza.pizza.toppings || []
      }
    })),
    combos: cart.combos || [],
  };
});

const initialState: CartSlice = {
  isCartOpened: false,
  cart: {
    products: [],
    pizzas: [],
    combos: [],
    totalCount: 0,
    totalPrice: 0,
  },
  loading: false,
  error: "",
};

export const cartSlice = createSlice({
  name: "cart",
  initialState,
  reducers: {
    setIsCartOpened: (state) => {
      state.isCartOpened = !state.isCartOpened;
    },
    clearCart: (state) => {
      state.cart = {
        products: [],
        pizzas: [],
        combos: [],
        totalCount: 0,
        totalPrice: 0,
      };
    },
    addNewProduct: (state, action: PayloadAction<ProductVariant>) => {
      state.cart.totalCount += 1;
      state.cart.totalPrice += action.payload.price;
      for (let i = 0; i < state.cart.products.length; i++) {
        if (
          JSON.stringify(state.cart.products[i].product) ===
          JSON.stringify(action.payload)
        ) {
          state.cart.products[i].count += 1;
          return;
        }
      }

      state.cart.products.push({
        product: action.payload,
        count: 1,
      });
    },
    addNewPizza: (state, action: PayloadAction<PizzaVariant>) => {
      state.cart.totalCount += 1;
      state.cart.totalPrice += action.payload.price;
      
      if (!state.cart.pizzas) {
        state.cart.pizzas = [];
      }
  
      const normalizedPizza = {
        ...action.payload,
        toppings: action.payload.toppings || []
      };
    
      const existingPizzaIndex = state.cart.pizzas.findIndex(pizza => {
        return (
          pizza.pizza.pizzaId === normalizedPizza.pizzaId &&
          pizza.pizza.size === normalizedPizza.size &&
          pizza.pizza.dough === normalizedPizza.dough &&
          JSON.stringify(pizza.pizza.toppings?.map(t => t.id).sort()) ===
          JSON.stringify(normalizedPizza.toppings?.map(t => t.id).sort())
        );
      });
    
      if (existingPizzaIndex !== -1) {
        state.cart.pizzas[existingPizzaIndex].count += 1;
      } else {
        state.cart.pizzas.push({
          pizza: normalizedPizza,
          count: 1,
        });
      }
    },

    deleteCartProduct: (state, action: PayloadAction<ProductVariant>) => {
      state.cart.totalCount -= 1;
      state.cart.totalPrice -= action.payload.price;

      for (let i = 0; i < state.cart.products.length; i++) {
        if (
          JSON.stringify(state.cart.products[i].product) ===
          JSON.stringify(action.payload)
        ) {
          state.cart.products[i].count -= 1;
          return;
        }
      }
    },

    deleteCartPizza: (state, action: PayloadAction<PizzaVariant>) => {
      state.cart.totalCount -= 1;
      state.cart.totalPrice -= action.payload.price;

      for (let i = 0; i < state.cart.pizzas.length; i++) {
        if (JSON.stringify(state.cart.pizzas[i].pizza) === JSON.stringify(action.payload)) {
          state.cart.pizzas[i].count -= 1;
          return;
        }
      }
    },

    deleteProductCompletely: (state, action: PayloadAction<ProductInCart>) => {
      state.cart.totalCount -= action.payload.count;
      state.cart.totalPrice -=
        action.payload.count * action.payload.product.price;

      for (let i = 0; i < state.cart.products.length; i++) {
        if (
          JSON.stringify(state.cart.products[i].product) ===
          JSON.stringify(action.payload.product)
        ) {
          state.cart.products.splice(i, 1);
          return;
        }
      }
    },

    deletePizzaCompletely: (state, action: PayloadAction<PizzaInCart>) => {
      state.cart.totalCount -= action.payload.count;
      state.cart.totalPrice -=
        action.payload.count * action.payload.pizza.price;

      for (let i = 0; i < state.cart.pizzas.length; i++) {
        if (
          JSON.stringify(state.cart.pizzas[i].pizza) ===
          JSON.stringify(action.payload.pizza)
        ) {
          state.cart.pizzas.splice(i, 1);
          return;
        }
      }
    }
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchCart.pending, (state) => {
        state.loading = true;
      })
      .addCase(fetchCart.fulfilled, (state, action) => {
        state.loading = false;
        state.cart = action.payload;
      })
      .addCase(fetchCart.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error as string;
      });
  },
});

export const {
  setIsCartOpened,
  addNewProduct,
  addNewPizza,
  deleteCartProduct,
  deleteCartPizza,
  deleteProductCompletely,
  deletePizzaCompletely,
  clearCart,
} = cartSlice.actions;

export default cartSlice.reducer;
