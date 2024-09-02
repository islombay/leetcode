package main

import (
	"fmt"
	"strconv"
)

func main() {
	cases := []struct {
		name   string
		n      int
		result bool
	}{
		{
			name:   "1",
			n:      19,
			result: true,
		},
		{
			name:   "2",
			n:      2,
			result: false,
		},
		{
			name:   "3",
			n:      7,
			result: true,
		},
	}

	for _, c := range cases {
		if result := isHappy(c.n); result != c.result {
			fmt.Printf("- Test %s failed, expected %t, got %t\n", c.name, c.result, result)
		} else {
			fmt.Printf("Test %s passed, expected %t, got %t\n", c.name, c.result, result)
		}
	}
}

func isHappy(n int) bool {
	ns := strconv.Itoa(n)

	digits := map[rune]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}

	for i := 0; i < 100; i++ {
		n = 0
		for _, c := range ns {
			n += pow(digits[c], 2)
		}

		ns = strconv.Itoa(n)

		if n == 1 {
			break
		}
	}

	if ns == "1" {
		return true
	}

	return false
}

func pow(base, exponent int) int {
	r := 1
	for exponent > 0 {
		r *= base
		exponent--
	}

	return r
}
