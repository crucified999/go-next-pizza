// "use client";

// import { fetchCart, setIsCartOpened } from "../../store/cartSlice";
// import { useAppDispatch, useAppSelector } from "@/app/store";
// import { cn } from "@/shared/lib/utils";
// import { X } from "lucide-react";
// import React, { useEffect, useState } from "react";
// import { CartProduct } from "../cart-item/cart-product";
// import { CartPizza } from "../cart-item/cart-pizza";

// export const Cart = () => {
//   const dispatch = useAppDispatch();
//   const isCartOpened = useAppSelector((state) => state.cart.isCartOpened);
//   const cart = useAppSelector((state) => state.cart.cart);

//   const handleClick = () => {
//     dispatch(setIsCartOpened());
//   };

//   useEffect(() => {
//     dispatch(fetchCart());
//   }, [dispatch]);

//   useEffect(() => {
//     if (isCartOpened) {
//       document.body.style.overflow = "hidden";
//     } else {
//       document.body.style.overflow = "auto";
//     }
//   }, [isCartOpened]);

//   return (

//     <>
//       <div
//         onClick={handleClick}
//         className={cn(
//           isCartOpened && "fixed z-100 top-0 left-0 w-full h-full bg-black/50"
//         )}
//       ></div>
//       <div
//         className={cn(
//           "fixed flex flex-col top-0 right-0 bg-gray-100 w-125 h-full transition-transform duration-300 ease-in-out gap-10",
//           isCartOpened ? "translate-x-0 z-150" : "translate-x-full"
//         )}
//       >
//         {cart.totalCount > 0 && (
//           <span className="ml-5 mt-10 text-2xl font-bold">{`${cart.totalCount} товара на ${cart.totalPrice} ₽`}</span>
//         )}
//         <button
//           onClick={handleClick}
//           className="rotate-0 transition-transform duration-700 linear hover:rotate-180 absolute top-1/2 right-135 -translate-y-[50%] cursor-pointer"
//         >
//           <X width={50} height={50} color="white" />
//         </button>
//         {cart.products && cart.products.length > 0 && (
//           <ul className="flex flex-col gap-3 overflow-auto">
//             {cart.products.map((product) => {
//               return (
//                 <li key={product.product.title + product.product.size}>
//                   <CartProduct cartItem={product} />
//                 </li>
//               );
//             })}
//           </ul>
//         )}

//         {cart.pizzas && cart.pizzas.length > 0 && (
//           <ul className="flex flex-col gap-3 overflow-auto">
//             {cart.pizzas.map((pizza) => {
//               return (
//                 <li
//                   key={pizza.pizza.title + pizza.pizza.size + pizza.pizza.dough}
//                 >
//                   <CartPizza cartItem={pizza} />
//                 </li>
//               );
//             })}
//           </ul>
//         )}
//       </div>
//     </>)
//   // ) : (
//   //   <div className="h-full flex flex-col items-center justify-center gap-5">
//   //     <img src="https://cdn.dodostatic.net/pizza-site/dist/assets/5aa5dac99a832c62f3ef..svg" />
//   //     <h4 className="font-bold text-2xl">Пока тут пусто</h4>
//   //     <p>Добавьте пиццу. Или две! А мы доставим ваш заказ от 1 ₽</p>
//   //   </div>
//   // );
// };

"use client";

import { fetchCart, setIsCartOpened } from "../../store/cartSlice";
import { useAppDispatch, useAppSelector } from "@/app/store";
import { cn } from "@/shared/lib/utils";
import { ArrowRight, X } from "lucide-react";
import React, { useEffect, useState } from "react";
import { CartProduct } from "../cart-item/cart-product";
import { CartPizza } from "../cart-item/cart-pizza";
import { declineProductWord } from "../../lib/utils";
import Link from "next/link";

export const Cart = () => {
  const dispatch = useAppDispatch();
  const isCartOpened = useAppSelector((state) => state.cart.isCartOpened);
  const cart = useAppSelector((state) => state.cart.cart);

  const [showEmptyCart, setShowEmptyCart] = useState(false);

  const handleClick = () => {
    dispatch(setIsCartOpened());
  };

  useEffect(() => {
    dispatch(fetchCart());
  }, [dispatch]);

  useEffect(() => {
    if (isCartOpened) {
      document.body.style.overflow = "hidden";
    } else {
      document.body.style.overflow = "auto";
    }
  }, [isCartOpened]);

  const hasItems = () => {
    const hasProducts = cart.products && cart.products.length > 0;
    const hasPizzas = cart.pizzas && cart.pizzas.length > 0;
    return hasProducts || hasPizzas;
  };

  useEffect(() => {
    if (!hasItems()) {
      const timer = setTimeout(() => {
        setShowEmptyCart(true);
      }, 100);
      return () => clearTimeout(timer);
    } else {
      setShowEmptyCart(false);
    }
  }, [cart.products, cart.pizzas]);

  return (
    <>
      <div
        onClick={handleClick}
        className={cn(
          "fixed z-100 top-0 left-0 w-full h-full bg-black/50 transition-opacity duration-300 ease-in-out",
          isCartOpened
            ? "opacity-100 pointer-events-auto"
            : "opacity-0 pointer-events-none"
        )}
      ></div>
      <div
        className={cn(
          "fixed grid grid-rows-[auto_1fr_auto] top-0 right-0 bg-gray-100 w-125 h-full transition-transform duration-300 ease-in-out dark:bg-[#101113]",
          isCartOpened ? "translate-x-0 z-150" : "translate-x-full",
          showEmptyCart && "flex flex-col"
        )}
      >
        {isCartOpened && (
          <button
            onClick={handleClick}
            className="rotate-0 transition-transform duration-700 linear hover:rotate-180 absolute top-1/2 right-135 -translate-y-[50%] cursor-pointer"
          >
            <X width={50} height={50} color="white" />
          </button>
        )}

        {hasItems() ? (
          <>
            {cart.totalCount > 0 && (
              <span className="ml-5 my-10 text-2xl font-bold">{`${
                cart.totalCount
              } ${declineProductWord(cart.totalCount)} на ${
                cart.totalPrice
              } ₽`}</span>
            )}

            <div className="overflow-auto flex-flex-col">
              {cart.products && cart.products.length > 0 && (
                <ul className="flex flex-col gap-3 overflow-auto px-5 pb-5">
                  {cart.products.map((product) => (
                    <li key={product.product.title + product.product.size}>
                      <CartProduct cartItem={product} />
                    </li>
                  ))}
                </ul>
              )}

              {cart.pizzas && cart.pizzas.length > 0 && (
                <ul className="flex flex-col gap-3 overflow-auto px-5 pb-5">
                  {cart.pizzas.map((pizza) => (
                    <li
                      key={
                        pizza.pizza.title +
                        pizza.pizza.size +
                        pizza.pizza.dough +
                        pizza.pizza.toppings.reduce((acc, v) => acc + v, "")
                      }
                    >
                      <CartPizza cartItem={pizza} />
                    </li>
                  ))}
                </ul>
              )}
            </div>

            <div className="bg-white flex flex-col gap-5 p-5 justify-end dark:bg-[#0a0b0b]">
              <div className="flex justify-between">
                <span className="font-bold text-lg">Сумма заказа</span>
                <span className="font-bold">{cart.totalPrice} ₽</span>
              </div>
              <Link
                href="/order"
                className="rounded-4xl text-lg justify-center relative bg-orange-500 text-white flex px-2 py-3 cursor-pointer hover:bg-[#e55400] transition-colors duration-300 linear border-none"
              >
                <span>К оформлению заказа</span>
                <ArrowRight className="absolute top-1/2 -translate-y-1/2 right-2" />
              </Link>
            </div>
          </>
        ) : (
          showEmptyCart && (
            <div className="h-full flex flex-col items-center justify-center gap-5 px-5 dark:bg-[#101113]">
              <img
                src="https://cdn.dodostatic.net/pizza-site/dist/assets/5aa5dac99a832c62f3ef..svg"
                alt="Пустая корзина"
                className="w-48 h-48 object-contain"
              />
              <h4 className="font-bold text-2xl text-center">Пока тут пусто</h4>
              <p className="text-gray-600 text-center dark:text-gray-400">
                Добавьте пиццу. Или две! А мы доставим ваш заказ от 1 ₽
              </p>
            </div>
          )
        )}
      </div>
    </>
  );
};
