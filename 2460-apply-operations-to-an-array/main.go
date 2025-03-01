package main

func main() {
	applyOperations([]int{847, 847, 0, 0, 0, 399, 416, 416, 879, 879, 206, 206, 206, 272})
}

func applyOperations(nums []int) []int {
	zeroIdx := -1

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 && zeroIdx == -1 {
			zeroIdx = i
			continue
		}

		if i+1 < len(nums) && nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}

		if zeroIdx != -1 && nums[i] != 0 {
			nums[zeroIdx] = nums[i]
			zeroIdx++
			nums[i] = 0
		}
	}

	return nums
}
