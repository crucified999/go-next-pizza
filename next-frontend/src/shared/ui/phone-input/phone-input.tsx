import React, {
  useState,
  useCallback,
  KeyboardEvent,
  ClipboardEvent,
  ChangeEvent,
  RefObject,
} from "react";

interface PhoneInputProps {
  value?: string;
  onChange?: (value: string) => void;
  placeholder?: string;
  label?: string;
  id?: string;
  className?: string;
  disabled?: boolean;
  required?: boolean;
  ref?: RefObject<HTMLInputElement | null>;
}

export const PhoneInput: React.FC<PhoneInputProps> = ({
  value,
  onChange,
  placeholder = "+7 999 999-99-99",
  label = "Номер телефона",
  id = "phone-input",
  className = "",
  disabled = false,
  required = false,
  ref = null
}) => {
  const [internalPhone, setInternalPhone] = useState<string>("+7 ");

  const phoneValue = value !== undefined ? value : internalPhone;

  const formatPhoneNumber = useCallback((value: string): string => {
    const cleaned: string = value.replace(/\D/g, "");
    let formatted: string = "+7 ";

    if (cleaned.length > 1) {
      const numbers: string = cleaned.slice(1);

      if (numbers.length > 0) {
        formatted += numbers.slice(0, 3);
      }
      if (numbers.length > 3) {
        formatted += " " + numbers.slice(3, 6);
      }
      if (numbers.length > 6) {
        formatted += "-" + numbers.slice(6, 8);
      }
      if (numbers.length > 8) {
        formatted += "-" + numbers.slice(8, 10);
      }
    }

    return formatted;
  }, []);

  const handlePhoneChange = useCallback(
    (e: ChangeEvent<HTMLInputElement>): void => {
      const input: string = e.target.value;

      if (input.length < 3) {
        const newValue = "+7 ";
        setInternalPhone(newValue);
        onChange?.(newValue);
        return;
      }

      const formatted: string = formatPhoneNumber(input);

      sessionStorage.setItem("phone-number", formatted);

      setInternalPhone(formatted);
      onChange?.(formatted);
    },

    [formatPhoneNumber, onChange]
  );

  const handleKeyDown = useCallback(
    (e: KeyboardEvent<HTMLInputElement>): void => {
      const allowedKeys: number[] = [8, 9, 13, 27, 46];

      if (e.ctrlKey && [65, 67, 86, 88].includes(e.keyCode)) {
        return;
      }

      if ([37, 38, 39, 40, 36, 35].includes(e.keyCode)) {
        return;
      }

      if (
        !allowedKeys.includes(e.keyCode) &&
        (e.keyCode < 48 || e.keyCode > 57)
      ) {
        e.preventDefault();
      }
    },
    []
  );

  const handlePaste = useCallback(
    (e: ClipboardEvent<HTMLInputElement>): void => {
      e.preventDefault();
      const pastedData: string = e.clipboardData.getData("text");
      const numbersOnly: string = pastedData.replace(/\D/g, "");

      if (numbersOnly) {
        const newValue: string = "+7 " + numbersOnly;
        const formatted: string = formatPhoneNumber(newValue);
        setInternalPhone(formatted);
        onChange?.(formatted);
      }
    },
    [formatPhoneNumber, onChange]
  );

  return (
    <div className={`flex flex-col gap-2 ${className}`}>
      {label && (
        <label htmlFor={id} className="text-sm">
          {label}
          {required && <span className="required-asterisk">*</span>}
        </label>
      )}
      <input
        id={id}
        type="tel"
        value={phoneValue}
        onChange={handlePhoneChange}
        onKeyDown={handleKeyDown}
        onPaste={handlePaste}
        placeholder={placeholder}
        disabled={disabled}
        required={required}
        maxLength={17}
        className="font-[600] p-3 rounded-3xl focus:outline-orange-500 bg-gray-100"
        ref={ref}
      />
    </div>
  );
};

export default PhoneInput;
