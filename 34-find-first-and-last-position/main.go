package main

import "fmt"

func main() {
	cases := []struct {
		name   string
		nums   []int
		target int
		result []int
	}{
		{
			name:   "#1",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			result: []int{3, 4},
		},
		{
			name:   "#2",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			result: []int{-1, -1},
		},
		{
			name:   "#3",
			nums:   []int{},
			target: 8,
			result: []int{-1, -1},
		},
	}

Out:
	for _, c := range cases {
		result := searchRange(c.nums, c.target)
		for i := 0; i < len(result); i++ {
			if result[i] != c.result[i] {
				fmt.Printf("Test %s failed. Result[i]: %d, c.result[i]: %d, index: %d", c.name, result[i], c.result[i], i)
				goto Out
			}
		}
		fmt.Printf("Test %s passed\n", c.name)
	}
}

func searchRange(nums []int, target int) []int {
	var (
		left  = 0
		right = len(nums) - 1
		start = -1
		end   = -1
	)

	if len(nums) == 0 {
		return []int{start, end}
	}

	// starting position
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			start = mid
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}

	if start == -1 {
		return []int{-1, -1}
	}

	left = start
	right = len(nums) - 1

	// end position
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			end = mid
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}

	return []int{start, end}
}
