package main

import "sort"

func maximumBeauty(nums []int, k int) int {
	if len(nums) == 1 {
		return 1
	}

	sort.Ints(nums)
	maxLen, r := 0, 1

	for l := range nums {
		for r < len(nums) && nums[r]-nums[l] <= 2*k {
			r++
		}
		maxLen = max(maxLen, r-l)
	}
	return maxLen
}
