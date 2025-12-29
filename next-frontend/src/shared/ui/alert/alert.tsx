"use client";

import { useEffect, useState } from "react";
import { CheckCircle } from "lucide-react";
import { useAppDispatch } from "@/app/store";
import { hideNotification, showNotification } from "./alertSlice";
import { cn } from "@/shared/lib/utils";

type AlertProps = {
  message: string;
  isVisible: boolean;
  duration?: number;
};

export const Alert = ({ message, isVisible, duration = 3000 }: AlertProps) => {
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(showNotification({ message }));

    if (isVisible) {
      const timer = setTimeout(() => {
        dispatch(hideNotification());
      }, duration);

      return () => clearTimeout(timer);
    }
  }, [isVisible, duration]);

  return (
    <div className={cn("absolute top-13 right-0 animate-fade")}>
      <div className="flex liquid-glass-element items-center gap-2 text-white px-4 py-3 rounded-lg shadow-lg text-sm">
        <CheckCircle size={15} />
        <span className="font-medium">{message}</span>
      </div>
    </div>
  );
};

