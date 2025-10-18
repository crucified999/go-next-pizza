export const CategoryList = () => {
  const categories = [
    {
      id: 1,
      name: "Пиццы",
    },
    {
      id: 2,
      name: "Комбо",
    },
    {
      id: 3,
      name: "Закуски",
    },
    {
      id: 4,
      name: "Коктейли",
    },
    {
      id: 5,
      name: "Кофе",
    },
    {
      id: 6,
      name: "Напитки",
    },
    {
      id: 7,
      name: "Десерты",
    },
    {
      id: 8,
      name: "Соусы",
    },
  ];

  return (
    <div className="py-10 grid">
      <h2 className="font-[800] text-4xl leading-[100%]">Категории</h2>
      <ul className="flex gap-5 p-4 px-0 rounded-2xl">
        {categories.map((category) => (
          <li className="hover:text-[#FE5F00] cursor-pointer transition-colors duration-150" key={category.id}>{category.name}</li>
        ))}
      </ul>
    </div>
  );
};
