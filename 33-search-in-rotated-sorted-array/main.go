package main

import "fmt"

func main() {
	cases := []struct {
		name   string
		nums   []int
		target int
		result int
	}{
		{
			name:   "Success test 1",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			result: 4,
		},
		{
			name:   "With Failure 2",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			result: -1,
		},
		{
			name:   "Sorted 3",
			nums:   []int{0, 1, 2, 3, 4, 5},
			target: 3,
			result: 3,
		},
	}

	for _, c := range cases {
		result := search(c.nums, c.target)
		if result != c.result {
			fmt.Printf("Test failed for test %s\n", c.name)
		} else if result == c.result {
			fmt.Printf("Test (%s) finished correctly\n", c.name)
		}
	}
}

func search(nums []int, target int) int {
	var (
		left  = 0
		right = len(nums) - 1
	)
	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[left] {
			// left part is sorted
			if nums[mid] > target && nums[left] <= target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < nums[left] {
			// right part is sorted
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
