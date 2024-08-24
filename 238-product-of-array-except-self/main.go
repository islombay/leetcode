package main

import (
	"fmt"
	"time"
)

func main() {
	cases := []struct {
		name   string
		nums   []int
		result []int
	}{
		{
			name:   "1",
			nums:   []int{1, 2, 3, 4},
			result: []int{24, 12, 8, 6},
		},
		{
			name:   "2",
			nums:   []int{-1, 1, 0, -3, 3},
			result: []int{0, 0, 9, 0, 0},
		},
	}

out:
	for _, c := range cases {
		startTime := time.Now()
		result := productExceptSelf(c.nums)
		elapsed := time.Since(startTime)

		fmt.Println("Time taken:", elapsed.Seconds())
		for i := range result {
			if result[i] != c.result[i] {
				fmt.Printf("- Test %s failed, expected %v, got %v\n", c.name, c.result, result)
				continue out
			}
		}
		fmt.Printf("Test %s passed, expected %v, got %v\n", c.name, c.result, result)
	}
}
func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	prefix := 1
	for i := 0; i < n; i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}
	return result
}
