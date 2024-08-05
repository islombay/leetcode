package main

import (
	"fmt"
	"math"
)

func main() {
	t1 := -1
	fmt.Println(reverse(t1))

	t1 = -123
	fmt.Println(reverse(t1))

	t1 = 120
	fmt.Println(reverse(t1))
}

func reverse(x int) int {
	var (
		currentVal       int
		previousWholeVal int = -1
		divider          int = 10
		reversed         []int
		reversedInt      int
		first            bool = true
	)

	if x == 0 {
		return 0
	} else if x == -1 {
		return x
	}
	for {
		currentVal = (x % divider)

		if first && currentVal == 0 {
			divider *= 10
			continue
		} else if first && currentVal != 0 {
			first = false
		}

		if currentVal == x && x == previousWholeVal {
			fmt.Println("broke here", currentVal, previousWholeVal)
			break
		}
		previousWholeVal = currentVal
		currentVal /= (divider / 10)

		reversed = append(reversed, currentVal)

		divider *= 10
	}

	fmt.Println(reversed)
	divider = len(reversed)
	for i, e := range reversed {
		reversedInt += e * intPow(10, divider-(i+1))
	}

	if reversedInt > math.MaxInt32 || reversedInt < math.MinInt32 {
		return 0
	}
	return reversedInt
}

func intPow(base, exponent int) int {
	result := 1
	for exponent > 0 {
		if exponent%2 == 1 {
			result *= base
		}
		base *= base
		exponent /= 2
	}
	return result
}
