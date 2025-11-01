import { Pizza } from "@/entities/product/model/types";
import { ToppingsList } from "@/entities/product/ui/toppings-list/toppings-list";
import { useToppings } from "@/shared/lib/hooks/useToppings";
import { ChevronLeft } from "lucide-react";

type ComboChoosePizzaProps = {
  pizza: Pizza;
  onReplace: () => void;
};

export const ComboChoosePizza = ({
  pizza,
  onReplace,
}: ComboChoosePizzaProps) => {
  const { selectedToppings, toggleTopping } = useToppings([]);
  return (
    <div className="flex flex-col overflow-auto px-5">
      <div className="flex items-center gap-5">
        <ChevronLeft
          className="cursor-pointer transition-transform duration-150 linear hover:scale-105 shadow-lg rounded-full"
          width={48}
          height={48}
          onClick={onReplace}
        />
        <h2 className="text-4xl font-[700]">Меняйте по вкусу</h2>
      </div>
      <div className="flex flex-col mt-2">
        <h3 className="text-xl font-[900]">Можно удалить</h3>
        <div className="flex flex-wrap gap-2 mb-5 mt-5">
          {pizza.ingredients.map((ingredient) =>
            ingredient.replacable ? (
              <button
                key={ingredient.id}
                className="bg-orange-500/15 text-orange-700 text-sm p-2 rounded-xl cursor-pointer"
              >
                {ingredient.title.toLowerCase()}
              </button>
            ) : null
          )}
        </div>
        <hr />
      </div>
      <div>
        <ToppingsList
          toppings={pizza.toppings || []}
          selectedToppings={selectedToppings}
          onToggleTopping={(toppingId: number) => {
            toggleTopping(toppingId);
          }}
        />
      </div>
    </div>
  );
};
