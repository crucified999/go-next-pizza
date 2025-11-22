'use client';

import { useAppDispatch } from "@/app/store";
import { setCurrentCategory } from "@/entities/category/store/categorySlice";
import { checkUserAuth } from "@/entities/user/store/userSlice";
import { PersonalDataForm } from "@/entities/user/ui/personal-data-form";
import { Footer } from "@/shared/ui/footer";
import { Header } from "@/shared/ui/header";
import { PostHeader } from "@/shared/ui/post-header";
import { useEffect } from "react";

export const ProfilePage = () => {
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(checkUserAuth());
    dispatch(setCurrentCategory(""));
  }, [dispatch]);

  return (
    <>
      <Header />
      <PostHeader />
      <PersonalDataForm />
      <Footer />
    </>
  );
}