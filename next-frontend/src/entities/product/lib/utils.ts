export const toppingsToMask = (toppingsIds: number[]): number => {
  let mask = 0;

  toppingsIds.forEach((tid) => {
    mask |= 1 << (tid - 1);
  });

  return mask;
}

export const maskToToppings = (mask: number): number[] => {
  const toppings = [];

  for (let i = 0; i < 64; i++) {
    toppings.push(i + 1);
  }

  return toppings;
}

export const convertDough = (doughType: number): string => {
  return doughType === 1 ? "традиционное" : "тонкое"
}

// func (pv *PizzaVariant) ToppingsToMask(toppingIDs []int) int {
// 	var mask int = 0
// 	for _, id := range toppingIDs {
// 			if id > 0 && id <= 64 {
// 					mask |= 1 << (id - 1)
// 			}
// 	}
// 	return mask
// }

// func (pv *PizzaVariant) MaskToToppings(mask int) []int {
// 	var toppings []int
// 	for i := 0; i < 64; i++ {
// 			if mask&(1<<i) != 0 {
// 					toppings = append(toppings, i+1)
// 			}
// 	}
// 	return toppings
// }