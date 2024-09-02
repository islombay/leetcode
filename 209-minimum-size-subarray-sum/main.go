package main

import (
	"fmt"
	"math"
)

func main() {
	cases := []struct {
		name   string
		nums   []int
		target int
		result int
	}{
		{
			name:   "1",
			nums:   []int{2, 3, 1, 2, 4, 3},
			target: 7,
			result: 2,
		},
		{
			name:   "2",
			nums:   []int{1, 4, 4},
			target: 4,
			result: 1,
		},
		{
			name:   "3",
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			target: 11,
			result: 0,
		},
	}

	for _, c := range cases {
		if res := minSubArrayLen(c.target, c.nums); res != c.result {
			fmt.Printf("- Test %s failed, expected %d, got %d\n", c.name, c.result, res)
		} else {
			fmt.Printf("Test %s passed, expected %d, got %d\n", c.name, c.result, res)
		}
	}
}

func minSubArrayLen(target int, nums []int) int {
	var (
		start      = 0
		currentSum = 0
		minLength  = math.MaxInt32
	)

	for end := 0; end < len(nums); end++ {
		currentSum += nums[end]

		for currentSum >= target {
			minLength = min(minLength, end-start+1)
			currentSum -= nums[start]
			start++
		}
	}

	if minLength == math.MaxInt32 {
		return 0
	}

	return minLength
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
