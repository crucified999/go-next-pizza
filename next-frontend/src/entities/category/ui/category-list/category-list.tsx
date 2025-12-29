"use client";

import { useAppDispatch, useAppSelector } from "@/app/store";
import {
  fetchCategories,
  setCurrentCategoryManually,
} from "../../store/categorySlice";
import { useEffect } from "react";
import { cn } from "@/shared/lib/utils";

export const CategoryList = () => {
  const categories = useAppSelector((state) => state.categories.categories).filter((cat) => cat.title !== 'Комбо');
  const currentCategory = useAppSelector(
    (state) => state.categories.currentCategory
  );
  const dispatch = useAppDispatch();

  const handleActiveCategory = (category: string) => {
    dispatch(setCurrentCategoryManually(category));
    localStorage.setItem("category", category);
  };

  useEffect(() => {
    dispatch(fetchCategories());
  }, [dispatch]);

  return (
    <div className="grid items-center">
      <ul className="flex gap-5 px-0 rounded-2x">
        {categories.map((category) => (
          <li
            className={cn(
              "hover:text-[#FE5F00] cursor-pointer transition-colors duration-150",
              currentCategory === category.title
                ? "text-[#FE5F00]"
                : "text-black dark:text-white"
            )}
            key={category.id}
            onClick={() => handleActiveCategory(category.title)}
          >
            <a onClick={() => {
              document.querySelector(`#${category.title}`)?.scrollIntoView({ behavior: 'smooth' })
            }}>{category.title}</a>
          </li>
        ))}
      </ul>
    </div>
  );
};
