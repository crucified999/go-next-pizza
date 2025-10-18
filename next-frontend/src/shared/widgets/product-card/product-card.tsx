import { Button } from '@/shared/ui/button';
import React from 'react';

type ProductCardProps = {
  id: number;
  title: string;
  description: string;
  price: number;
  image: string;
  choosable: boolean;
}
export const ProductCard: React.FC<ProductCardProps> = ({ id, title, description, price, image, choosable }) => {
  return (
    <article className="flex flex-col gap-2 shrink-1 max-w-[220px] h-full">
      <img src={image} alt={title} className="cursor-pointer max-w-[218px] max-h-[218px] transition-all duration-300 linear translate-y-0 hover:translate-y-[5px]" />
      <h3 className="text-xl font-bold">{title}</h3>
      <p className="text-sm opacity-80 flex-grow">{description}</p>
      
      <div className="mt-auto flex justify-between items-center">
        <span className="text-lg font-[600]">{choosable ? `от ${price}` : price} ₽</span>
        <Button variant="outline" className="border-none">
          {choosable ? <span>Выбрать</span> : <span>В корзину</span>}
        </Button>
      </div>
    </article>
  );
}