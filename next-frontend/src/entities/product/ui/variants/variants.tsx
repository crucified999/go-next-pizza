"use client";

import { cn } from "@/shared/lib/utils";
import { useEffect, useRef, useState } from "react";

type Option = {
  value: string;
  available: boolean;
}

type VariantsProps = {
  options: Option[];
  value?: number;
  onChange?: (index: number) => void;
};

export const Variants = ({ options, value, onChange }: VariantsProps) => {
  const [internalIndex, setInternalIndex] = useState(0);
  const activeIndex = value ?? internalIndex;
  const containerRef = useRef<HTMLUListElement | null>(null);
  const itemRefs = useRef<HTMLLIElement[]>([]);
  const [indicator, setIndicator] = useState<{ width: number; left: number }>({
    width: 0,
    left: 0,
  });

  const recalcIndicator = (index: number) => {
    const el = itemRefs.current[index];
    if (!el || !containerRef.current) return;
  
    const left = el.offsetLeft;
    const width = el.offsetWidth;
    setIndicator({ width, left });
  };

  useEffect(() => {
    recalcIndicator(activeIndex);
  }, [activeIndex, options.length]);

  useEffect(() => {
    const onResize = () => recalcIndicator(activeIndex);
    window.addEventListener("resize", onResize);
    const id = window.setTimeout(() => recalcIndicator(activeIndex), 0);
    return () => {
      window.removeEventListener("resize", onResize);
      window.clearTimeout(id);
    };
  }, [activeIndex]);

  return (
    <ul
      ref={containerRef}
      className="relative cursor-pointer flex bg-[#ECECEC] rounded-xl py-1 px-2 mt-2 gap-2 w-full justify-around dark:bg-[#101113]"
    >
      <div
        className="pointer-events-none absolute inset-y-1 bg-white dark:bg-[#101113] dark:border-1 dark:border-white rounded-xl transition-all duration-300 ease-in-out z-0"
        style={{ width: `${indicator.width}px`, left: `${indicator.left}px` }}
      />

      {options.map((option, idx) => (
        <li
          key={option.value}
          ref={(el) => {
            if (el) itemRefs.current[idx] = el;
          }}
          onClick={() => {
            if (option.available) {
              onChange ? onChange(idx) : setInternalIndex(idx);
            }
            
          }}
          className={cn("relative z-[1] w-full text-center font-[700] text-xs border-none text-black dark:text-white px-3 py-1.5 rounded-md select-none", !option.available && "opacity-50 cursor-not-allowed")}
        >
          {option.value}
        </li>
      ))}
    </ul>
  );
};
