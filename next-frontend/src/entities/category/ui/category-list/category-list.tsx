"use client";

import { useAppDispatch, useAppSelector } from "@/app/store";
import {
  fetchCategories,
  setCurrentCategoryManually,
} from "../../store/categorySlice";
import { useEffect } from "react";
import { cn } from "@/shared/lib/utils";
import { useRestoreCategory } from "@/shared/lib/hooks/useRestoreCategory";

export const CategoryList = () => {
  const categories = useAppSelector((state) => state.categories.categories);
  const currentCategory = useAppSelector(
    (state) => state.categories.currentCategory
  );
  const dispatch = useAppDispatch();

  // useRestoreCategory();

  const handleActiveCategory = (category: string) => {
    dispatch(setCurrentCategoryManually(category));
    sessionStorage.setItem("category", category);
  };

  useEffect(() => {
    dispatch(fetchCategories());
  }, [dispatch]);

  return (
    <div className="grid items-center">
      <ul className="flex gap-5 px-0 rounded-2xl">
        {categories.map((category) => (
          <li
            className={cn(
              "hover:text-[#FE5F00] cursor-pointer transition-colors duration-150",
              currentCategory === category.title
                ? "text-[#FE5F00]"
                : "text-black"
            )}
            key={category.id}
            onClick={() => handleActiveCategory(category.title)}
          >
            <a href={`#${category.title}`}>{category.title}</a>
          </li>
        ))}
      </ul>
    </div>
  );
};
