"use client";

import { Modal } from "@/shared/ui/modal";
import { AuthForm } from "../auth-form";

export const AuthModal = () => {
  return (
    <Modal className="w-[430px] h-[350px] p-5">
      <div>
        <div className="flex flex-col justify-center items-center">
          <h3 className="font-bold text-2xl">Укажите телефон</h3>
          <p className="text-black/50">Чтобы войти в профиль</p>
        </div>
        <AuthForm />
        <FormFooter />
      </div>
    </Modal>
  );
};

export const FormFooter = () => {
  return (
    <div className="flex flex-col mt-7 justify-center items-center">
      <p>Продолжая, вы соглашаетесь с условиями наших</p>
      <a
        className="text-orange-500"
        href="https://dodopizza.ru/moscow/1may/legal"
        target="_blank"
      >
        юридических документов
      </a>
    </div>
  );
};
