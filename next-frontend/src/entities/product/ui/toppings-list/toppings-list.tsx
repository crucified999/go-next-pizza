import { Topping } from "../../model/types";
import { ToppingCard } from "../topping-card";

type ToppingsListProps = {
  toppings: Topping[];
  selectedToppings: number[];
  onToggleTopping: (toppingId: number) => void;
};

export const ToppingsList = ({
  toppings,
  selectedToppings,
  onToggleTopping,
}: ToppingsListProps) => {
  return (
    <div className="mt-10">
      <h3 className="text-xl font-bold">Добавить по вкусу</h3>
      <ul className="grid grid-cols-3 gap-2 py-5 max-h-[300px]">
        {toppings.map((topping) => (
          <li key={topping.id}>
            <ToppingCard
              topping={topping}
              isSelected={selectedToppings.includes(topping.id)}
              onToggle={() => onToggleTopping(topping.id)}
            />
          </li>
        ))}
      </ul>
    </div>
  );
};
