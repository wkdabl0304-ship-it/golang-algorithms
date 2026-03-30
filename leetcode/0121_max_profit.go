package main

import "fmt"

func maxProfit(prices []int) int {
	minPrice := prices[0]
	res := 0
	for _, v := range prices {
		if v < minPrice {
			minPrice = v
		}
		if v-minPrice > res {
			res = v - minPrice
		}
	}
	return res
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
