import { PhoneInput } from "@/shared/ui/phone-input";
import { SyntheticEvent, useEffect, useState } from "react";
import { sendCode } from "../../lib/api";
import { normalizePhoneNumber } from "../../lib/utils";
import Link from "next/link";
import { FormButton } from "@/shared/ui/form-button";
import { useRouter } from "next/navigation";

export const AuthForm = () => {
  const [phoneValue, setPhoneValue] = useState<string>("+7 ");
  const router = useRouter();

  const isValid = phoneValue.replace(/\D/g, "").length === 11;

  const handleSubmit = async (e: SyntheticEvent) => {
    e.preventDefault();

    if (isValid) {
      try {
        const normalizedPhone = normalizePhoneNumber(phoneValue);
        // Сохраняем номер телефона в sessionStorage перед отправкой
        sessionStorage.setItem("phone-number", phoneValue);
        await sendCode(normalizedPhone);
        // Переходим на страницу SMS только после успешной отправки
        router.replace("/sms");
      } catch (error) {
        console.error("Ошибка при отправке кода:", error);
        // Можно добавить уведомление пользователю об ошибке
        alert("Не удалось отправить код. Попробуйте еще раз.");
      }
    }
  };

  useEffect(() => {
    const savedPhone = sessionStorage.getItem("phone-number");

    if (savedPhone) {
      setPhoneValue(savedPhone);
    }
  });

  return (
    <form className="mt-5 w-full flex flex-col gap-2" onSubmit={handleSubmit}>
      <PhoneInput value={phoneValue} onChange={setPhoneValue} />

      <FormButton disabled={!isValid} text="Выслать код" />
    </form>
  );
};
