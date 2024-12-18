package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	res := addDigits(234567873456765434)
	elapsed := time.Now().Sub(start)
	fmt.Printf("Result: %d, took: %s\n", res, elapsed)
}

func addDigits(num int) int {
	if num < 10 {
		return num
	}

	runeToNum := map[rune]int{
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

	strNum := strconv.Itoa(num)
	sum := 0

	for len(strNum) > 1 {
		sum = 0
		for _, v := range strNum {
			sum += runeToNum[v]
		}
		strNum = strconv.Itoa(sum)
	}
	return sum
}
