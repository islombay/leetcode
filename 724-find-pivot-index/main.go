package main

func pivotIndex(nums []int) int {
	sumRight := make([]int, len(nums))
	sumRight[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		sumRight[i] = sumRight[i+1] + nums[i]
	}

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sumRight[i] == sum {
			return i
		}
	}

	return -1
}
