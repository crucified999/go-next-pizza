import React, { memo } from "react";

type FormButtonProps = {
  onClick?: () => void;
  disabled: boolean;
  text: string;
};

export const FormButton: React.FC<FormButtonProps> = memo(({ onClick, disabled, text }) => {
  return (
    <button
      onClick={onClick}
      type="submit"
      disabled={disabled}
      className="w-full cursor-pointer hover:bg-orange-500/80 transition-colors duration-300 linear bg-orange-500 p-3 rounded-3xl text-white disabled:opacity-50 disabled:cursor-not-allowed disabled:hover:bg-orange-500"
    >
      {text}
    </button>
  );
});
