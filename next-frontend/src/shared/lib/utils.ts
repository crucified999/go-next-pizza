import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}


export function formatPhoneFromE164(e164Phone: string): string {
  const cleanPhone = e164Phone.replace(/[^\d+]/g, '');

  if (!cleanPhone.startsWith('+7') && !cleanPhone.startsWith('7') && !cleanPhone.startsWith('8')) {
    return cleanPhone;
  }

  let normalizedPhone = cleanPhone;
  
  if (cleanPhone.startsWith('8')) {
    normalizedPhone = '+7' + cleanPhone.slice(1);
  } else if (cleanPhone.startsWith('7') && !cleanPhone.startsWith('+7')) {
    normalizedPhone = '+' + cleanPhone;
  } else if (cleanPhone.startsWith('+7')) {
    normalizedPhone = cleanPhone;
  }
  
  if (normalizedPhone.length !== 12) {
    return normalizedPhone;
  }
  
  const countryCode = normalizedPhone.slice(0, 2); 
  const operatorCode = normalizedPhone.slice(2, 5);
  const firstPart = normalizedPhone.slice(5, 8);   
  const secondPart = normalizedPhone.slice(8, 10); 
  const thirdPart = normalizedPhone.slice(10, 12); 
  
  return `${countryCode} ${operatorCode} ${firstPart} ${secondPart} ${thirdPart}`;
}