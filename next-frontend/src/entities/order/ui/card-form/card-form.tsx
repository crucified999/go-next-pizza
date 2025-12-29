'use client';

import React, { useState, useEffect, useRef } from 'react';
import { cn } from '@/shared/lib/utils';
import { CardIcon } from '@/shared/ui/card-icon';
import { detectCardType, isValidCardNumber } from '../../lib/utils';
import Link from 'next/link';
import { useRouter } from 'next/navigation';

interface CardFormProps {
  onSubmit?: (cardData: CardData) => void;
  isLoading?: boolean;
  className?: string;
  showSaveOption?: boolean;
  amount?: number;
}

export interface CardData {
  number: string;
  expiry: string;
  cvv: string;
  cardholder: string;
  saveCard?: boolean;
}

export const CardForm: React.FC<CardFormProps> = ({
  onSubmit,
  isLoading = false,
  className = '',
  showSaveOption = true,
  amount,
}) => {
  const router = useRouter();

  const [formData, setFormData] = useState<CardData>({
    number: '',
    expiry: '',
    cvv: '',
    cardholder: '',
    saveCard: false,
  });
  
  const [errors, setErrors] = useState<Record<string, string>>({});
  const [focusedField, setFocusedField] = useState<string | null>(null);
  const cardInfo = detectCardType(formData.number);

  const cvvRef = useRef<HTMLInputElement>(null);
  
  useEffect(() => {
    if (formData.expiry.length === 5 && cvvRef.current) {
      cvvRef.current.focus();
    }
  }, [formData.expiry]);
  
  const formatCardNumber = (value: string): string => {
    const cleaned = value.replace(/\s/g, '').replace(/\D/g, '');
    const groups = cleaned.match(/.{1,4}/g);
    return groups ? groups.join(' ') : cleaned;
  };
  
  const formatExpiry = (value: string): string => {
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
  };
  
  const validateField = (name: string, value: string): string => {
    switch (name) {
      case 'number':
        if (!value.trim()) return 'Введите номер карты';
        const cleanNumber = value.replace(/\s/g, '');
        if (cleanNumber.length < 16) return 'Номер карты должен содержать 16 цифр';
        if (!isValidCardNumber(cleanNumber)) return 'Неверный номер карты';
        return '';
        
      case 'expiry':
        if (!value.trim()) return 'Введите срок действия';
        if (!/^\d{2}\/\d{2}$/.test(value)) return 'Формат: ММ/ГГ';
        
        const [month, year] = value.split('/').map(Number);
        const currentDate = new Date();
        const currentYear = currentDate.getFullYear() % 100;
        const currentMonth = currentDate.getMonth() + 1;
        
        if (month < 1 || month > 12) return 'Неверный месяц';
        if (year < currentYear || (year === currentYear && month < currentMonth)) {
          return 'Карта просрочена';
        }
        return '';
        
      case 'cvv':
        if (!value.trim()) return 'Введите CVV';
        if (!/^\d{3,4}$/.test(value)) return 'CVV должен содержать 3-4 цифры';
        return '';
        
      case 'cardholder':
        if (!value.trim()) return 'Введите имя держателя карты';
        if (value.length < 2) return 'Имя должно содержать минимум 2 символа';
        return '';
        
      default:
        return '';
    }
  };
  
  const handleChange = (field: keyof CardData, value: string | boolean) => {
    let formattedValue = value;
    
    if (typeof value === 'string') {
      switch (field) {
        case 'number':
          formattedValue = formatCardNumber(value).slice(0, 19);
          break;
        case 'expiry':
          formattedValue = formatExpiry(value).slice(0, 5);
          break;
        case 'cvv':
          formattedValue = value.replace(/\D/g, '').slice(0, 4);
          break;
        case 'cardholder':
          formattedValue = value.toUpperCase().slice(0, 50);
          break;
      }
    }
    
    setFormData(prev => ({ ...prev, [field]: formattedValue }));
    
    if (typeof value === 'string' && errors[field]) {
      const error = validateField(field, formattedValue as string);
      setErrors(prev => ({
        ...prev,
        [field]: error,
      }));
    }
  };
  
  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    
    const newErrors: Record<string, string> = {};
    Object.keys(formData).forEach(key => {
      if (key === 'saveCard') return;
      
      const error = validateField(key, formData[key as keyof CardData] as string);
      if (error) newErrors[key] = error;
    });
    
    setErrors(newErrors);
    
    if (Object.keys(newErrors).length === 0) {
      onSubmit?.(formData);

      router.replace('/success');
    }

    
  };
  
  const handleFocus = (field: string) => {
    setFocusedField(field);
  };
  
  const handleBlur = () => {
    setFocusedField(null);
  };
  
  const getCardMaskSimple = () => {
    if (!formData.number) return '•••• •••• •••• ••••';
    
    const cleanNumber = formData.number.replace(/\s/g, '');
    const totalDigits = 16;
    const visibleDigits = Math.min(cleanNumber.length, totalDigits);
    
    let result = '';
    for (let i = 0; i < totalDigits; i++) {
      if (i < visibleDigits) {
        result += cleanNumber[i];
      } else {
        result += '•';
      }
      
      if ((i + 1) % 4 === 0 && i < totalDigits - 1) {
        result += ' ';
      }
    }
    
    return result;
  };

  return (
    <div className={cn('w-full max-w-md', className)}>
      <div className="mb-6">
        <div className="relative bg-gradient-to-r from-orange-500 to-orange-600 rounded-xl p-6 text-white shadow-lg">
          <div className="absolute top-4 left-4 w-10 h-8 bg-yellow-400 rounded-md flex items-center justify-center">
            <div className="w-6 h-4 bg-gradient-to-r from-yellow-300 to-yellow-500 rounded-sm" />
          </div>
          
          <div className="absolute top-4 right-4">
            <CardIcon 
              cardNumber={formData.number} 
              className="text-white"
              size="lg"
            />
          </div>
          
          <div className="mt-12 mb-6">
            <div className="text-2xl font-mono tracking-wider">
              {getCardMaskSimple()}
            </div>
            <div className="text-sm text-white/70 mt-1">
              Номер карты
            </div>
          </div>
          <div className="flex justify-between items-end">
            <div>
              <div className="text-sm mb-1">ДЕРЖАТЕЛЬ КАРТЫ</div>
              <div className="font-medium tracking-wider">
                {formData.cardholder || 'ВАШЕ ИМЯ'}
              </div>
            </div>
            
            <div className="text-right">
              <div className="text-sm mb-1">СРОК ДЕЙСТВИЯ</div>
              <div className="font-medium tracking-wider">
                {formData.expiry || 'ММ/ГГ'}
              </div>
            </div>
          </div>
        </div>
      </div>
      
      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">
            Номер карты
          </label>
          <div className="relative">
            <input
              type="text"
              value={formData.number}
              onChange={(e) => handleChange('number', e.target.value)}
              onFocus={() => handleFocus('number')}
              onBlur={handleBlur}
              placeholder="0000 0000 0000 0000"
              inputMode="numeric"
              className={cn(
                'w-full px-4 py-3 pl-12 border rounded-lg focus:outline-none focus:ring-2 transition-all',
                'font-mono text-lg tracking-wider',
                errors.number 
                  ? 'border-red-500 focus:border-red-500 focus:ring-red-200' 
                  : focusedField === 'number'
                    ? 'border-orange-500 focus:border-orange-500 focus:ring-orange-200'
                    : 'border-gray-300 focus:border-orange-500 focus:ring-orange-200'
              )}
              maxLength={19}
            />
            
            <div className="absolute left-3 top-1/2 transform -translate-y-1/2">
              <CardIcon cardNumber={formData.number} size="sm" />
            </div>
            
            <button
              type="button"
              className="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-orange-500 transition-colors"
              title="Сканировать карту"
            >
              <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} 
                      d="M3 9l2-2m0 0l7-7 7 7M5 9v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
              </svg>
            </button>
          </div>
          {errors.number && (
            <p className="mt-1 text-sm text-red-600">{errors.number}</p>
          )}
        </div>
        
        <div className="grid grid-cols-2 gap-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Срок действия
            </label>
            <input
              type="text"
              value={formData.expiry}
              onChange={(e) => handleChange('expiry', e.target.value)}
              onFocus={() => handleFocus('expiry')}
              onBlur={handleBlur}
              placeholder="ММ/ГГ"
              inputMode="numeric"
              className={cn(
                'w-full px-4 py-3 border rounded-lg focus:outline-none focus:ring-2 transition-all',
                errors.expiry 
                  ? 'border-red-500 focus:border-red-500 focus:ring-red-200' 
                  : focusedField === 'expiry'
                    ? 'border-orange-500 focus:border-orange-500 focus:ring-orange-200'
                    : 'border-gray-300 focus:border-orange-500 focus:ring-orange-200'
              )}
              maxLength={5}
            />
            {errors.expiry && (
              <p className="mt-1 text-sm text-red-600">{errors.expiry}</p>
            )}
          </div>
          
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              CVV
              <button
                type="button"
                className="ml-1 text-gray-400 hover:text-orange-500 transition-colors"
                title="3 цифры на обратной стороне карты"
              >
                <svg className="w-4 h-4 inline" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} 
                        d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </button>
            </label>
            <div className="relative">
              <input
                ref={cvvRef}
                type="password"
                value={formData.cvv}
                onChange={(e) => handleChange('cvv', e.target.value)}
                onFocus={() => handleFocus('cvv')}
                onBlur={handleBlur}
                placeholder="123"
                inputMode="numeric"
                maxLength={4}
                className={cn(
                  'w-full px-4 py-3 border rounded-lg focus:outline-none focus:ring-2 transition-all',
                  'font-mono',
                  errors.cvv 
                    ? 'border-red-500 focus:border-red-500 focus:ring-red-200' 
                    : focusedField === 'cvv'
                      ? 'border-orange-500 focus:border-orange-500 focus:ring-orange-200'
                      : 'border-gray-300 focus:border-orange-500 focus:ring-orange-200'
                )}
              />
    
              <div className="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400">
                <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} 
                        d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                </svg>
              </div>
            </div>
            {errors.cvv && (
              <p className="mt-1 text-sm text-red-600">{errors.cvv}</p>
            )}
          </div>
        </div>
        
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">
            Имя держателя карты
          </label>
          <input
            type="text"
            value={formData.cardholder}
            onChange={(e) => handleChange('cardholder', e.target.value)}
            onFocus={() => handleFocus('cardholder')}
            onBlur={handleBlur}
            placeholder="IVAN IVANOV"
            className={cn(
              'w-full px-4 py-3 border rounded-lg focus:outline-none focus:ring-2 transition-all',
              'uppercase',
              errors.cardholder 
                ? 'border-red-500 focus:border-red-500 focus:ring-red-200' 
                : focusedField === 'cardholder'
                  ? 'border-orange-500 focus:border-orange-500 focus:ring-orange-200'
                  : 'border-gray-300 focus:border-orange-500 focus:ring-orange-200'
            )}
          />
          {errors.cardholder && (
            <p className="mt-1 text-sm text-red-600">{errors.cardholder}</p>
          )}
        </div>
        
        {showSaveOption && (
          <div className="flex items-center">
            <input
              type="checkbox"
              id="saveCard"
              checked={formData.saveCard}
              onChange={(e) => handleChange('saveCard', e.target.checked)}
              className="h-4 w-4 focus:ring-orange-500 border-gray-300 rounded accent-orange-500"
            />
            <label htmlFor="saveCard" className="ml-2 block text-sm text-gray-700">
              Сохранить карту для будущих платежей
            </label>
          </div>
        )}
        
        <button
          type="submit"
          disabled={isLoading}
          className={cn(
            'w-full py-3 px-4 font-medium transition-all rounded-3xl',
            'flex items-center justify-center gap-2',
            isLoading
              ? 'bg-gray-400 cursor-not-allowed text-white'
              : 'bg-orange-500 hover:bg-orange-600 transition-colors duration-300 text-white shadow-md hover:shadow-lg cursor-pointer'
          )}
        >
          {isLoading ? (
            <>
              <svg className="animate-spin h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              Обработка...
            </>
          ) : (
            <>
              <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} 
                      d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z" />
              </svg>
              {amount ? `Оплатить ${amount} ₽` : 'Добавить карту'}
            </>
          )}
        </button>
      
      </form>
    </div>
  );
};