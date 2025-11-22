"use client";

import { User } from "lucide-react";
import { Button } from "../../ui/button";
import Link from "next/link";
import { useAppSelector } from "@/app/store";

export const ProfileButton = () => {
  const isAuth = useAppSelector((state) => state.user.isAuth);

  return (
    <Link href={isAuth ? "/profile" : "/auth"}>
      <Button variant="outline" className="p-5">
        <User width={12} />
        <span>{isAuth ? "Профиль" : "Войти"}</span>
      </Button>
    </Link>
  );
};
