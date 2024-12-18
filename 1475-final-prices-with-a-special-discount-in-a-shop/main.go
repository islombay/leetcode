package main

func finalPrices(prices []int) []int {
	if len(prices) == 1 {
		return prices
	}

	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[i] >= prices[j] {
				prices[i] -= prices[j]
				break
			}
		}
	}

	return prices
}
