import { useEffect } from "react";
import { useAppDispatch } from "@/app/store";
import { setCurrentCategoryManually } from "@/entities/category/store/categorySlice";

export const useRestoreCategory = () => {
  const dispatch = useAppDispatch();

  useEffect(() => {
    if (typeof window !== "undefined") {
      const savedCategory = sessionStorage.getItem("category");
      
      if (savedCategory) {
        dispatch(setCurrentCategoryManually(savedCategory));
      }
    }
  }, [dispatch]);
};