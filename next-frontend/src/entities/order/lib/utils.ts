export type CardType = 
  | 'visa' 
  | 'mastercard' 
  | 'amex' 
  | 'discover' 
  | 'diners' 
  | 'jcb' 
  | 'unionpay' 
  | 'maestro' 
  | 'mir' 
  | 'unknown';

export function detectCardType(cardNumber: string): { type: CardType; name: string } {
  const cleaned = cardNumber.replace(/\s/g, '');
  
  if (!cleaned) {
    return { type: 'unknown', name: 'Карта' };
  }
  
  if (cleaned.startsWith('4')) {
    return { type: 'visa', name: 'Visa' };
  }
  
  const firstTwo = parseInt(cleaned.substring(0, 2));
  if (firstTwo >= 51 && firstTwo <= 55) {
    return { type: 'mastercard', name: 'MasterCard' };
  }
  
  const firstFour = parseInt(cleaned.substring(0, 4));
  if (firstFour >= 2221 && firstFour <= 2720) {
    return { type: 'mastercard', name: 'MasterCard' };
  }
  
  if (cleaned.startsWith('34') || cleaned.startsWith('37')) {
    return { type: 'amex', name: 'American Express' };
  }

  if (cleaned.startsWith('2200') || cleaned.startsWith('2204')) {
    return { type: 'mir', name: 'МИР' };
  }
  
  return { type: 'unknown', name: 'Карта' };
}

export function isValidCardNumber(cardNumber: string): boolean {
  const cleaned = cardNumber.replace(/\s/g, '');
  
  if (!/^\d+$/.test(cleaned)) {
    return false;
  }
  
  let sum = 0;
  let isEven = false;
  
  for (let i = cleaned.length - 1; i >= 0; i--) {
    let digit = parseInt(cleaned.charAt(i));
    
    if (isEven) {
      digit *= 2;
      if (digit > 9) {
        digit -= 9;
      }
    }
    
    sum += digit;
    isEven = !isEven;
  }
  
  return sum % 10 === 0;
}

export function formatCardNumber(value: string): string {
  const cleaned = value.replace(/\s/g, '').replace(/\D/g, '');
  const groups = cleaned.match(/.{1,4}/g);
  return groups ? groups.join(' ') : cleaned;
}

export function formatExpiry(value: string): string {
  const cleaned = value.replace(/\D/g, '');
  
  if (cleaned.length >= 2) {
    const month = cleaned.substring(0, 2);
    const year = cleaned.substring(2, 4);
    
    let formatted = month;
    if (month.length === 2 && cleaned.length > 2) {
      formatted += `/${year}`;
    }
    
    return formatted;
  }
  
  return cleaned;
}

export const formatDateTimeExact = (isoString: string): string => {
  if (!isoString) return '';
  
  try {
    const date = new Date(isoString);
    
    if (isNaN(date.getTime())) {
      return isoString;
    }
    
    const months = [
      'янв.', 'февр.', 'мар.', 'апр.', 'мая', 'июн.',
      'июл.', 'авг.', 'сент.', 'окт.', 'нояб.', 'дек.'
    ];
    
    const day = date.getDate();
    const month = months[date.getMonth()];
    const year = date.getFullYear();
    const hours = date.getHours().toString().padStart(2, '0');
    const minutes = date.getMinutes().toString().padStart(2, '0');
    
    return `${day} ${month} ${year} г., ${hours}:${minutes}`;
  } catch (error) {
    console.error('Error formatting date:', error);
    return isoString;
  }
};