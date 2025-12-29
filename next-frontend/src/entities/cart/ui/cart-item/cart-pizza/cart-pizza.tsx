import { useAppDispatch } from "@/app/store";
import { addPizza, deletePizza } from "@/entities/cart/lib/api";
import { PizzaInCart } from "@/entities/cart/model/types";
import { addNewPizza, deleteCartPizza, deletePizzaCompletely, deleteProductCompletely } from "@/entities/cart/store/cartSlice";
import { toppingsToMask } from "@/entities/product/lib/utils";
import { X } from "lucide-react";
import { useEffect, useState } from "react";

type CartPizzaProps = {
  cartItem: PizzaInCart;
};

export const CartPizza: React.FC<CartPizzaProps> = ({ cartItem }) => {
  const dispatch = useAppDispatch();
  const [count, setCount] = useState(cartItem.count);

  const handleAddProduct = async () => {
    const addedPizza = await addPizza({
      pizzaId: cartItem.pizza.pizzaId,
      dough: cartItem.pizza.dough,
      size: cartItem.pizza.size,
      toppingsMask: toppingsToMask(cartItem.pizza.toppings.map((t) => t.id)),
    });

    dispatch(addNewPizza(addedPizza));
    setCount((prev) => prev + 1);
  };

  const handleDeletePizza = async () => {
    await deletePizza({
      pizzaId: cartItem.pizza.pizzaId,
      dough: cartItem.pizza.dough,
      size: cartItem.pizza.size,
      toppingsMask: toppingsToMask(cartItem.pizza.toppings.map((t) => t.id)),
      action: "delete",
    });
    dispatch(deleteCartPizza(cartItem.pizza));
    setCount((prev) => prev - 1);
  };

  const handleDeleteCompletely = async () => {
    await deletePizza({
      pizzaId: cartItem.pizza.pizzaId,
      dough: cartItem.pizza.dough,
      size: cartItem.pizza.size,
      toppingsMask: toppingsToMask(cartItem.pizza.toppings.map((t) => t.id)),
      action: "delete-completely",
    });
    dispatch(deletePizzaCompletely(cartItem));
    setCount(0);
  };

  useEffect(() => {
    setCount(cartItem.count);
  }, [cartItem.count]);

  return (
    count > 0 && (
      <article className="bg-white w-auto p-5 relative rounded-2xl dark:bg-[#0f1113] dark:border-1 dark:border-white">
        <button
          onClick={handleDeleteCompletely}
          className="absolute top-5 right-10 cursor-pointer"
        >
          <X width={20} height={20} />
        </button>
        <div className="grid grid-cols-[auto_1fr] gap-5 mb-3">
          <img
            className="w-16 h-16"
            src={cartItem.pizza.image}
            alt={cartItem.pizza.title}
          />
          <div>
            <h4 className="text-xl font-[600]">{cartItem.pizza.title}</h4>
            <span className="text-sm text-black/50 dark:text-white">
              {cartItem.pizza.size},{" "}
              {cartItem.pizza.dough === 1
                ? "традиционное"
                : "тонкое" + " тесто"}
            </span>

            <p className="text-xs text-black/70 font-[600]">
              {cartItem.pizza.toppings.length > 0 && "+ "}
              {cartItem.pizza.toppings.map((t, index) => {
                return (
                  <span key={t.title}>
                    {t.title.toLowerCase() +
                      `${
                        index !== cartItem.pizza.toppings.length - 1
                          ? ", "
                          : " "
                      }`}
                  </span>
                );
              })}
            </p>
          </div>
        </div>
        <hr />
        <div className="flex justify-between mt-3">
          <span className="font-[600]">
            {cartItem.pizza.price + (cartItem.pizza.toppings &&
              cartItem.pizza.toppings?.reduce((acc, t) => acc + t.price, 0))}{" "}
            ₽
          </span>
          <div className="flex gap-2 justify-center items-center">
            <button className="text-orange-500 cursor-pointer hover:text-[#e55400] transition-colors duration-150 linear text-[15px]">
              Изменить
            </button>
            <div className="flex items-center justify-around gap-6 bg-gray-100 dark:bg-[#000000] border-none rounded-2xl py-1 px-3">
              <button
                onClick={handleDeletePizza}
                className="cursor-pointer text-xl"
              >
                –
              </button>
              <span>{count}</span>
              <button
                onClick={handleAddProduct}
                className="cursor-pointer text-xl"
              >
                +
              </button>
            </div>
          </div>
        </div>
      </article>
    )
  );
};
