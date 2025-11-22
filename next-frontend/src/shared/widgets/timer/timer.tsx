import React, { useEffect, useImperativeHandle, useRef, useState } from "react";
import { TIMER } from "./constants";
import { cn } from "@/shared/lib/utils";

type TimerProps = {
  onTimerChange?: (timer: number) => void;
};

export const Timer = React.forwardRef<{ reset: () => void }, TimerProps>(
  ({ onTimerChange }, ref) => {
    const [timer, setTimer] = useState(TIMER);
    const prevTimerRef = useRef(timer);
    const onTimerChangeRef = useRef(onTimerChange);

    useEffect(() => {
      onTimerChangeRef.current = onTimerChange;
    }, [onTimerChange]);

    useEffect(() => {
      onTimerChangeRef.current?.(TIMER);
    }, []);

    useEffect(() => {
      const interval = setInterval(() => {
        setTimer((prev) => {
          const newValue = prev <= 1 ? 0 : prev - 1;
          if (prevTimerRef.current > 0 && newValue === 0) {
            onTimerChangeRef.current?.(0);
          } else if (prevTimerRef.current === 0 && newValue > 0) {
            onTimerChangeRef.current?.(newValue);
          }
          prevTimerRef.current = newValue;
          return newValue;
        });
      }, 1000);

      return () => clearInterval(interval);
    }, []);

    useImperativeHandle(
      ref,
      () => ({
        reset: () => {
          setTimer(TIMER);
          prevTimerRef.current = TIMER;
          onTimerChangeRef.current?.(TIMER);
        },
      }),
      []
    );

    const timerDisplay = timer > 0 ? `Получить новый код через ${timer}` : null;

    if (!timerDisplay) return null;

    return (
      <span className={cn("opacity-0", timer > 0 && "opacity-100")}>
        {timerDisplay}
      </span>
    );
  }
);

Timer.displayName = "Timer";
