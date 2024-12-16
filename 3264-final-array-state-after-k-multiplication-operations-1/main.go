package main

func getFinalState(nums []int, k int, multiplier int) []int {
	n := len(nums)
	idx := 0

	for ; k > 0; k-- {
		idx = 0

		for j := 1; j < n; j++ {
			if nums[idx] > nums[j] {
				idx = j
			}
		}
		nums[idx] *= multiplier
	}
	return nums
}
