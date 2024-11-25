package main

import "fmt"

func main() {
	cases := []struct {
		name   string
		prices []int
		result int
	}{
		{
			name:   "1",
			prices: []int{7, 1, 5, 3, 6, 4},
			result: 5,
		},
		{
			name:   "2",
			prices: []int{7, 6, 4, 3, 1},
			result: 0,
		},
		{
			name:   "3",
			prices: []int{7, 6, 5, 3, 4, 5},
			result: 2,
		},
		{
			name:   "4",
			prices: []int{7, 3, 5, 2, 2},
			result: 2,
		},
	}

	for _, c := range cases {
		result := maxProfit2(c.prices)
		if c.result != result {
			fmt.Printf("- Test %s failed, expected %d, got %d\n", c.name, c.result, result)
		} else {
			fmt.Printf("Test %s passed, expected %d, got %d\n", c.name, c.result, result)
		}
	}
}

func maxProfit2(prices []int) int {
	minPrice := prices[0]
	maxPrice := 0

	for _, v := range prices {
		diff := v - minPrice
		if diff > maxPrice {
			maxPrice = diff
		}

		if v < minPrice {
			minPrice = v
		}
	}
	return maxPrice
}

func maxProfit(prices []int) int {
	var (
		minPrice  = prices[0]
		maxProfit = 0
	)

	for _, e := range prices {
		if e-minPrice > maxProfit {
			maxProfit = e - minPrice
		}
		if e < minPrice {
			minPrice = e
		}
	}

	return maxProfit
}
