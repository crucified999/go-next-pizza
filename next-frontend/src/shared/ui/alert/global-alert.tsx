"use client";

import { useAppDispatch, useAppSelector } from "@/app/store";
import { Alert } from "./alert";
import { hideNotification } from "./alertSlice";

export const GlobalAlert = () => {
  // const dispatch = useAppDispatch();
  const { isVisible, message } = useAppSelector((state) => state.notification);

  return (
    <Alert
      message={message}
      isVisible={isVisible}
    />
  );
};