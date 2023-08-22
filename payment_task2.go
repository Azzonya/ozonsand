package main

import (
	"fmt"
	"sort"
)

func main() {
	var t int
	fmt.Scan(&t) // Количество наборов входных данных

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n) // Количество купленных товаров в данном наборе

		counts := make(map[int]int)
		for j := 0; j < n; j++ {
			var price int
			fmt.Scan(&price) // Считываем цену товара
			counts[price]++
		}

		prices := make([]int, 0)
		for price := range counts {
			prices = append(prices, price)
		}

		sort.Ints(prices)

		totalCost := 0
		for _, price := range prices {
			count := counts[price]

			// Применяем акцию, если товаров больше 3
			paidCount := count - (count / 3)
			totalCost += paidCount * price
		}

		fmt.Println(totalCost)
	}
}
