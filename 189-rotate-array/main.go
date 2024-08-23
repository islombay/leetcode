package main

import "fmt"

func main() {
	cases := []struct {
		name   string
		nums   []int
		k      int
		result []int
	}{
		{
			name:   "1",
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			k:      3,
			result: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name:   "2",
			nums:   []int{-1, -100, 3, 99},
			k:      2,
			result: []int{3, 99, -1, -100},
		},
	}
Out:
	for _, c := range cases {
		rotate(c.nums, c.k)
		if len(c.nums) == len(c.result) {
			for i, v := range c.nums {
				if v != c.result[i] {
					fmt.Printf("Test %s failed, expected %v, got %v\n", c.name, c.result, c.nums)
					continue Out
				}
			}
			fmt.Printf("Test %s passed\n", c.name)
		} else {
			fmt.Printf("Test %s failed, expected %v, got %v\n", c.name, c.result, c.nums)
		}
	}
}

func rotate(nums []int, k int) {
	k = k % len(nums)

	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)

	for i, e := range numsCopy {
		nums[(i+k)%len(nums)] = e
	}
}
