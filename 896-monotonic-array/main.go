package main

func main() {

}

func isMonotonic(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	// 1 1 2
	movingTo := 0
	// 0 not found
	// 1 down
	// 2 up

	for i := 1; i < len(nums); i++ {
		if movingTo == 0 {
			if nums[i-1] < nums[i] {
				movingTo = 2
			} else if nums[i-1] > nums[i] {
				movingTo = 1
			}
			continue
		}
		if nums[i-1] < nums[i] && movingTo == 1 {
			return false
		}
		if nums[i-1] > nums[i] && movingTo == 2 {
			return false
		}
	}
	return true
}
