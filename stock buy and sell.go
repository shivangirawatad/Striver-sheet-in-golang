package main

import "fmt"

func maxProfit(prices []int) int {
	min := prices[0]
	max := prices[0]
	profit := 0

	for i := range prices {
		if prices[i] > max {
			max = prices[i]
		}
		if prices[i] < min {
			min = prices[i]
			max = min
		}
		if max-min > profit {
			profit = max - min
		}
	}

	return profit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	result := maxProfit(prices)
	fmt.Println(result)
}
