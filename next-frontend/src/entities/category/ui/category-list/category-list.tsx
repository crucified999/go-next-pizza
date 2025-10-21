"use client";

import { useAppDispatch, useAppSelector } from "@/app/store";
import { fetchCategories } from "../../store/categorySlice";
import { useEffect } from "react";

export const CategoryList = () => {
  const categories = useAppSelector((state) => state.categories.categories);
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(fetchCategories());
  }, [dispatch]);

  return (
    <div className="pt-10 grid">
      <h2 className="font-[800] text-4xl leading-[100%]">Категории</h2>
      <ul className="flex gap-5 pt-4 px-0 rounded-2xl">
        {categories.map((category) => (
          <li
            className="hover:text-[#FE5F00] cursor-pointer transition-colors duration-150"
            key={category.id}
          >
            {category.title}
          </li>
        ))}
      </ul>
    </div>
  );
};
