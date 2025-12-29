"use client";

import { useAppSelector } from "@/app/store";
import { cn } from "@/shared/lib/utils";
import React, { ChangeEvent, useEffect, useState } from "react";

type DataInputProps = {
  label: string;
  value: string;
  canChange: boolean;
  inputClassname?: string;
  onSave?: (id: number, value: string) => void;
};

export const DataInput: React.FC<DataInputProps> = ({
  label,
  value,
  canChange,
  inputClassname,
  onSave,
}) => {
  const userId = useAppSelector((state) => state.user.id);
  const [isFocused, setIsFocused] = useState(false);
  const [inputValue, setInputValue] = useState(value);
  const [isInChangeMode, setIsInChangeMode] = useState(false);

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    setInputValue(e.target.value);
  };

  const handleSave = () => {
    onSave?.(userId, inputValue);
    setIsInChangeMode(false);
  };

  const handleCancel = () => {
    setInputValue(value);
    setIsInChangeMode(false);
  }

  useEffect(() => {
    setInputValue(value);
  }, [value]);

  return (
    <div className="flex gap-3 items-center">
      <div
        className={cn(
          !canChange && "cursor-not-allowed",
          !isInChangeMode && "cursor-not-allowed",
          "flex flex-col gap-1"
        )}
      >
        <label className="text-sm" htmlFor={label}>
          {label}
        </label>
        <div className={cn("bg-gray-100 dark:bg-[#101113] p-3 rounded-2xl border-1 border-transparent transition-colors duration-100 linear", isFocused && "border-orange-500")}>
          <input
            onFocus={() => setIsFocused(true)}
            onBlur={() => setIsFocused(false)}
            onChange={handleChange}
            required
            className={cn(
              "dark:text-gray-200",
              !canChange && "pointer-events-none text-black/50 dark:text-gray-400",
              !isInChangeMode && "pointer-events-none text-black/50",
              "border-none focus:outline-none",
              inputClassname,
            )}
            type="text"
            name={label}
            value={inputValue}
          />
          {isInChangeMode ? (
            <button
              onClick={handleSave}
              disabled={inputValue === value}
              type="submit"
              className={cn(
                !canChange && "opacity-0 pointer-events-none",
                "cursor-pointer border-none text-orange-500 text-[15px]",
                inputValue === value && "text-black/50 cursor-disabled",
              )}
            >
              Сохранить
            </button>
          ) : (
            <button
              onClick={() => setIsInChangeMode(true)}
              type="button"
              className={cn(
                !canChange && "opacity-0 pointer-events-none",
                "cursor-pointer border-none text-orange-500 text-[15px]"
              )}
            >
              Изменить
            </button>
          )}
        </div>
      </div>

      {!canChange ||
        (isInChangeMode && (
          <button
            onClick={handleCancel}
            className="mt-7 cursor-pointer text-orange-500 text-lg"
          >
            Отменить
          </button>
        ))}
    </div>
  );
};
