"use client";

import { X } from "lucide-react";
import { cn } from "@/shared/lib/utils";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

interface ModalProps {
  children: React.ReactNode;
  className?: string;
}

export const Modal = ({ children, className }: ModalProps) => {
  const router = useRouter();

  const handleClose = () => {
    router.back();
  };

  useEffect(() => {
    document.body.style.overflow = "hidden";
    return () => {
      document.body.style.overflow = "auto";
    };
  }, []);

  return (
    <div className={cn("fixed inset-0 flex justify-center items-center z-100 bg-black/30")}>
      <div
        className={cn(
          "relative w-[1000px] h-[700px] max-h-4xl bg-white rounded-xl dark:bg-[#101113]",
          className
        )}
      >
        <button
          className="transition-transform dark:bg-[#101113] duration-300 linear cursor-pointer absolute top-5 -right-20 p-4 rounded-full bg-white hover:scale-110"
          onClick={handleClose}
        >
          <X className="w-6 h-6" />
        </button>
        {children}
      </div>
    </div>
  );
};
