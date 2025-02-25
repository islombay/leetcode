package main

import "sort"

func countPairs(nums []int, target int) int {
	sort.Ints(nums)

	ans, left, right := 0, 0, len(nums)-1

	for left < right {
		if nums[right]+nums[left] < target {
			ans += right - left
			left++
		} else {
			right--
		}
	}
	return ans
}
