package main

import (
	"fmt"
	"math"
	"unicode"
)

func main() {
	a := "+-12"
	// fmt.Println(myAtoi(a))

	a = "     -042333"
	// fmt.Println(myAtoi(a))

	a = "10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459"
	fmt.Println(myAtoi(a))

	a = "     -042"
	// fmt.Println(myAtoi(a))

	a = "9223372036854775808"
	// a = "18446744073709551617"
	// fmt.Println(myAtoi(a))

	a = "-91283472332"
	// fmt.Println("result", myAtoi(a))
}

func myAtoi(s string) int {
	var (
		isPositive bool  = true
		first      bool  = true
		digits     []int = []int{}
		res        int   = 0
	)

	nums := map[rune]int{
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

	for _, r := range s {
		if first {
			if r == ' ' {
				continue
			} else if r == '+' || r == '-' {
				if r == '-' {
					isPositive = false
				}
				first = false
			} else if !unicode.IsDigit(r) {
				return 0
			} else if unicode.IsDigit(r) {
				if r == '0' {
					continue
				} else {
					digits = append(digits, nums[r])
				}
				first = false
			}
		} else {
			if !unicode.IsDigit(r) {
				break
			} else if unicode.IsDigit(r) {
				digits = append(digits, nums[r])
			}
		}
	}

	for i, n := range digits {
		adding := n * pow(10, len(digits)-i-1)
		fmt.Println(adding, len(digits))
		if res+adding > math.MaxInt32 && isPositive || res > res+adding || n != 0 && adding == 0 {
			res = math.MaxInt32
			break
		} else if (adding+res)*-1 < math.MinInt32 && !isPositive || res > res+adding {
			res = math.MinInt32
			isPositive = true
			break
		}
		res += adding
	}

	if !isPositive {
		res *= -1
	}

	return res
}

func pow(n, p int) int {
	var r int = 1
	for i := 0; i < p; i++ {
		r *= n
	}
	return r
}
