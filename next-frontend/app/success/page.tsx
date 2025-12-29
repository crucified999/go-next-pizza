'use client';

import { useRouter } from 'next/navigation';
import { CheckCircle } from 'lucide-react';

export default function SuccessPage() {
  const router = useRouter();

  return (
    <div className="min-h-screen flex flex-col items-center justify-center p-4">
      <div className="text-center mb-8">
        <div className="w-24 h-24 bg-orange-100 rounded-full flex items-center justify-center mx-auto mb-6">
          <CheckCircle className="w-16 h-16 text-orange-500" />
        </div>
        
        <h1 className="text-2xl font-bold text-gray-900 mb-2">
          Спасибо за заказ!
        </h1>
        <p className="text-gray-600 mb-8">
          Мы свяжемся с вами для подтверждения
        </p>
      </div>

      <button
        onClick={() => router.replace('/')}
        className="px-8 py-3 bg-orange-500 text-white rounded-lg hover:bg-orange-600 transition-colors"
      >
        Вернуться на главную
      </button>
    </div>
  );
}