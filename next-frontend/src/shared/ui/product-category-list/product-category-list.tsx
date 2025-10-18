import React from 'react';
import { ProductCard } from '@/shared/widgets/product-card/product-card';

type ProductCategoryListProps = {
  category: string;
}

export const ProductCategoryList: React.FC<ProductCategoryListProps> = ({ category }) => {

  const products = [
    {
      id: 1,
      title: "Пицца 1",
      description: "Двойная порция классических колбасок, красный лук, томаты, соус барбекю, моцарелла, фирменный томатный соус",
      price: 100,
      image: "https://media.dodostatic.net/image/r:584x584/0198bf57bc517218ab93c762f4b0193e.avif",
      choosable: true,
    },
    {
      id: 2,
      title: "Пицца 1",
      description: "Двойная порция классических колбасок, красный лук, томаты, соус барбекю, моцарелла, фирменный томатный соус",
      price: 100,
      image: "https://media.dodostatic.net/image/r:584x584/0198bf57bc517218ab93c762f4b0193e.avif",
      choosable: true,
    },
    {
      id: 3,
      title: "Пицца 1",
      description: "Двойная порция классических колбасок, красный лук, томаты, соус барбекю, моцарелла, фирменный томатный соус",
      price: 100,
      image: "https://media.dodostatic.net/image/r:584x584/0198bf57bc517218ab93c762f4b0193e.avif",
      choosable: true,
    },
    {
      id: 4,
      title: "Пицца 1",
      description: "Двойная порция классических колбасок, красный лук, томаты, соус барбекю, моцарелла, фирменный томатный соус",
      price: 100,
      image: "https://media.dodostatic.net/image/r:584x584/0198bf57bc517218ab93c762f4b0193e.avif",
      choosable: true,
    },
    {
      id: 5,
      title: "Пицца 1",
      description: "Двойная порция классических колбасок, красный лук, томаты, соус барбекю, моцарелла, фирменный томатный соус",
      price: 100,
      image: "https://media.dodostatic.net/image/r:584x584/0198bf57bc517218ab93c762f4b0193e.avif",
      choosable: true,
    },
  ]; // TODO: Взять из стора

  return (
    <div className="py-15">
      <h2 className="font-[800] text-3xl leading-[100%]">{category}</h2>
      <ul className="grid grid-cols-4 gap-25 py-20">
        {products.map((product) => (
        <li key={product.id}>
          <ProductCard id={product.id} title={product.title} description={product.description} price={product.price} image={product.image} choosable={product.choosable} />
        </li>
      ))}
      </ul>
    </div>
  );
};