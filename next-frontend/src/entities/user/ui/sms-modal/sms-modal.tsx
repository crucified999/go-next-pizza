"use client";

import { Modal } from "@/shared/ui/modal";
import { SMSForm } from "../sms-form";
import { FormFooter } from "../auth-modal/auth-modal";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";

export const SMSModal = () => {
  const [phone, setPhone] = useState("");
  const router = useRouter();

  useEffect(() => {
    const savedPhone = sessionStorage.getItem("phone-number");

    if (savedPhone) {
      setPhone(savedPhone);
    }
  });

  return (
    <Modal className="w-[430px] h-[350px] p-5">
      <div className="flex flex-col justify-center items-center mb-5">
        <h3 className="font-bold text-2xl">Введите код</h3>
        <div className="flex gap-2">
          <p className="text-black/50">из СМС на номер {phone}</p>

          <button onClick={() => router.replace("/auth")} className="bg-transparent text-md border-none text-orange-500 cursor-pointer">
            Изменить
          </button>
        </div>
      </div>
      <SMSForm />
      <FormFooter />
    </Modal>
  );
};
