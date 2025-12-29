'use client';

import { useSticky } from "@/shared/lib/hooks/useSticky";
import { useRef } from "react";
import { CategoryList } from "@/entities/category/ui/category-list";
import { CartButton } from "@/shared/widgets/cart-button";

export const PostHeader = () => {
  const stickyRef = useRef<HTMLDivElement>(null!);
  const isSticky = useSticky(stickyRef);  

  return (
    <div
        ref={stickyRef}
        className={`dark:bg-[#101113] flex justify-between py-2 my-10 sticky top-0 left-0 right-0 z-10 transition-all duration-300 ${
          isSticky ? "bg-gray-200 px-50" : "bg-white"
        }`}
        style={
          isSticky
            ? {
                width: "100vw",
                marginLeft: "calc(-50vw + 50%)",
              }
            : {}
        }
      >
        <div className="flex items-center gap-10">
          {isSticky && (
            <img src="/logo.png" alt="logo" className="w-[35px] h-[35px]" />
          )}
          <CategoryList />
        </div>

        <CartButton />
      </div>
  )
};