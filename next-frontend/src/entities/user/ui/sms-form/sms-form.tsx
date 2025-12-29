"use client";

import { FormButton } from "@/shared/ui/form-button";
import { InputOTP } from "@/shared/ui/input-otp";
import {
  InputOTPGroup,
  InputOTPSeparator,
  InputOTPSlot,
} from "@/shared/ui/input-otp/input-otp";
import { REGEXP_ONLY_DIGITS } from "input-otp";
import React, {
  SyntheticEvent,
  useCallback,
  useMemo,
  useRef,
  useState,
} from "react";
import { sendCode } from "../../lib/api";
import { SMS_VERIFICATION_FAILED } from "../../lib/constants";
import { normalizePhoneNumber } from "../../lib/utils";
import { Timer } from "@/shared/widgets/timer";
import { useAppDispatch } from "@/app/store";
import { verifyAuth } from "../../store/userSlice";
import { useRouter } from "next/navigation";
import { fetchCart } from "@/entities/cart/store/cartSlice";

export const SMSForm = () => {
  const [code, setCode] = useState("");
  const [shouldResetTimer, setShouldResetTimer] = useState(0);
  const [isTimerActive, setIsTimerActive] = useState(true);
  const timerRef = useRef<{ reset: () => void } | null>(null);
  const dispatch = useAppDispatch();
  const router = useRouter();

  const inputRef = useRef<HTMLInputElement>(null);

  const handleSubmit = useCallback(
    async (e: SyntheticEvent) => {
      e.preventDefault();

      const result = await dispatch(verifyAuth(inputRef.current!.value));

      if (verifyAuth.fulfilled.match(result)) {
        if (result.payload.verified) {
          localStorage.setItem("isLoggedIn", "true");
          router.push("/");
          dispatch(fetchCart());
        } else {
          setShouldResetTimer((prev) => prev + 1);
          setIsTimerActive(true);
        }
      } else {
        setShouldResetTimer((prev) => prev + 1);
        setIsTimerActive(true);
      }

    },
    [dispatch, router]
  );

  const handleNewClick = () => {
    const savedPhone = sessionStorage.getItem("phone-number");

    if (savedPhone) {
      sendCode(normalizePhoneNumber(savedPhone));
      setShouldResetTimer((prev) => prev + 1);
      setIsTimerActive(true);
    }
  };

  const handleTimerChange = useCallback((timer: number) => {
    setIsTimerActive(timer > 0);
  }, []);

  const isButtonDisabled = useMemo(() => code.length !== 6, [code.length]);

  return (
    <form
      onSubmit={handleSubmit}
      className="flex flex-col justify-center items-center gap-3"
    >
      <InputOTP
        onChange={() => setCode(inputRef.current!.value)}
        ref={inputRef}
        maxLength={6}
        pattern={REGEXP_ONLY_DIGITS}
        className="w-full"
      >
        <InputOTPGroup>
          <InputOTPSlot
            className="rounded-2xl h-15 w-10"
            index={0}
            autoFocus={true}
          />
          <InputOTPSeparator />
          <InputOTPSlot className="rounded-2xl h-15 w-10" index={1} />
          <InputOTPSeparator />
          <InputOTPSlot className="rounded-2xl h-15 w-10" index={2} />
          <InputOTPSeparator />
          <InputOTPSlot className="rounded-2xl h-15 w-10" index={3} />
          <InputOTPSeparator />
          <InputOTPSlot className="rounded-2xl h-15 w-10" index={4} />
          <InputOTPSeparator />
          <InputOTPSlot className="rounded-2xl h-15 w-10" index={5} />
        </InputOTPGroup>
      </InputOTP>

      <Timer
        ref={timerRef}
        key={shouldResetTimer}
        onTimerChange={handleTimerChange}
      />

      {isTimerActive ? (
        <FormButton disabled={isButtonDisabled} text="Подтвердить" />
      ) : (
        <button
          className="bg-orange-500 transition-colors duration-300 linear hover:bg-orange-500/80 text-white p-3 rounded-3xl w-full cursor-pointer"
          onClick={handleNewClick}
        >
          Получить новый код
        </button>
      )}
    </form>
  );
};
