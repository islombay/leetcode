package main

import (
	"fmt"
	"sort"
)

func main() {
	cases := []struct {
		name   string
		nums   []int
		result int
	}{
		{
			name:   "1",
			nums:   []int{3, 2, 3},
			result: 3,
		},
		{
			name:   "2",
			nums:   []int{2, 2, 1, 1, 1, 2, 2},
			result: 2,
		},
	}

	for _, c := range cases {
		result := majorityElementWithSort(c.nums)
		if result != c.result {
			fmt.Printf(" - Test %s failed, expected %d, got %d\n", c.name, c.result, result)
		} else {
			fmt.Printf("Test %s passed, expected %d, got %d\n", c.name, c.result, result)
		}
	}
}

func majorityElementWithSort(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func majorityElement(nums []int) int {
	var (
		counter    = make(map[int]int)
		maxElement = 0
		maxSize    = 0
	)

	for _, n := range nums {
		counter[n]++
		val, ok := counter[n]
		if ok && val > maxSize {
			maxSize = val
			maxElement = n
		}
	}
	return maxElement
}
