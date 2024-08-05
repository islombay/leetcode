package main

import "fmt"

func main() {
	a := []int{9}
	fmt.Println(plusOne(a))
}

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}

	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
