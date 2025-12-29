// components/ui/card-icon.tsx
import React from 'react';
import { cn } from '@/shared/lib/utils';
import { detectCardType } from '@/entities/order/lib/utils';

interface CardIconProps {
  cardNumber: string;
  className?: string;
  size?: 'sm' | 'md' | 'lg';
}

export const CardIcon: React.FC<CardIconProps> = ({
  cardNumber,
  className = '',
  size = 'md',
}) => {
  const cardInfo = detectCardType(cardNumber);
  
  const sizeClasses = {
    sm: 'w-6 h-6',
    md: 'w-8 h-8',
    lg: 'w-12 h-12',
  };
  
  const getIcon = () => {
    switch (cardInfo.type) {
      case 'visa':
        return (
          <svg className={sizeClasses[size]} viewBox="0 0 24 24" fill="none">
            <path d="M9.6 15.4H7.3l1.4-8.7h2.3l-1.4 8.7z" fill="#1A1F71"/>
            <path d="M15.2 6.7c-1.2 0-2.1.7-2.6 1.6l-3.4 6.5h2.5l.5-1.2h3.1l.3 1.2h2.2l-1.9-8.1h-1.7zm.3 5.1l.8-2.2.8 2.2h-1.6z" fill="#EB001B"/>
            <path d="M22.2 6.7c-.8 0-1.3.4-1.7 1l-4.9 6.5h2.6l.7-1.8h3.2l.4 1.8h2.2l-2.5-8.1h-1.5zm.6 5.1l1-2.7 1 2.7h-2z" fill="#F79E1B"/>
          </svg>
        );
        
      case 'mastercard':
        return (
          <svg className={sizeClasses[size]} viewBox="0 0 24 24" fill="none">
            <circle cx="9" cy="12" r="7" fill="#EB001B"/>
            <circle cx="15" cy="12" r="7" fill="#F79E1B"/>
            <path d="M12 5a7 7 0 000 14 7 7 0 000-14z" fill="#FF5F00"/>
          </svg>
        );
        
      case 'mir':
        return (
          <svg className={sizeClasses[size]} viewBox="0 0 24 24" fill="none">
            <circle cx="12" cy="12" r="12" fill="#1E5A96"/>
            <path d="M15 9h-2v6h2V9zm-4 0H9v6h2V9zm-4 0H5v6h2V9z" fill="white"/>
          </svg>
        );
        
      case 'amex':
        return (
          <div className={cn('bg-blue-600 text-white rounded flex items-center justify-center font-bold', 
            sizeClasses[size], className)}>
            AMEX
          </div>
        );
        
      default:
        return (
          <div className={cn('bg-gray-200 rounded flex items-center justify-center', 
            sizeClasses[size], className)}>
            <svg className="w-3/4 h-3/4 text-gray-400" fill="currentColor" viewBox="0 0 20 20">
              <path fillRule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-6-3a2 2 0 11-4 0 2 2 0 014 0zm-2 4a5 5 0 00-4.546 2.916A5.986 5.986 0 005 10a5.986 5.986 0 00.454 2.084A5 5 0 0010 15a5 5 0 004.546-2.916A5.986 5.986 0 0015 10a5.986 5.986 0 00-.454-2.084A5 5 0 0010 7z" clipRule="evenodd"/>
            </svg>
          </div>
        );
    }
  };
  
  return (
    <div className={cn('inline-flex items-center', className)}>
      {getIcon()}
    </div>
  );
};