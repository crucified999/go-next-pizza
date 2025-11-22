"use client";

import { useEffect } from "react";
import { useAppDispatch } from "@/app/store";
import { checkUserAuth } from "@/entities/user/store/userSlice";

export const useCheckAuth = () => {
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(checkUserAuth());
  }, [dispatch]);
};


