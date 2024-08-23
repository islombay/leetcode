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
			result: 7,
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
		{
			name:   "5",
			prices: []int{1, 2, 3, 4, 5},
			result: 4,
		},
	}

	for _, c := range cases {
		result := maxProfit(c.prices)
		if c.result != result {
			fmt.Printf("- Test %s failed, expected %d, got %d\n", c.name, c.result, result)
		} else {
			fmt.Printf("Test %s passed, expected %d, got %d\n", c.name, c.result, result)
		}
	}
}
func maxProfit(prices []int) int {
	var (
		minPrice  = prices[0]
		maxProfit = 0
	)
	// []int{7, 6, 5, 3, 4, 5}

	for i, e := range prices {
		if e <= minPrice {
			minPrice = e
		} else if i+1 < len(prices) {
			if prices[i+1] < e {
				maxProfit += e - minPrice
				minPrice = prices[i+1]
			}
		} else {
			maxProfit += e - minPrice
		}
	}

	return maxProfit
}
