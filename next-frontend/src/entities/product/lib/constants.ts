const categoryMap = {
  "pizza": "Пиццы",
  "combo": "Комбо",
  "snack": "Закуски",
  "dessert": "Десерты",
  "coffee": "Кофе",
  "shake": "Коктейли",
  "drink": "Напитки",
  "sauce": "Соусы",
}

export const getCategoryName = (category: string) => {
  return categoryMap[category as keyof typeof categoryMap];
}