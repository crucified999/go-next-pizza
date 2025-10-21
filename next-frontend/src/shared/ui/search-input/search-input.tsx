import { Search } from "lucide-react";

export const SearchInput = () => {
  return (
    <div className="p-3 flex items-center gap-2 bg-[#F9F9F9] rounded-[15px] grow-1">
      <Search width={16} className="opacity-[0.5]" />
      <input type="text" placeholder="Поиск пиццы..." className="focus:outline-none w-full" />
    </div>
  );
};
