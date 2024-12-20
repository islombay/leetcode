package main

func sortColors(nums []int) {
	redCount := 0
	whiteCount := 0
	blueCount := 0
	idx := 0

	for _, n := range nums {
		if n == 0 {
			redCount++
		} else if n == 1 {
			whiteCount++
		} else {
			blueCount++
		}
	}

	for redCount > 0 {
		nums[idx] = 0
		idx++
		redCount--
	}
	for whiteCount > 0 {
		nums[idx] = 1
		idx++
		whiteCount--
	}
	for blueCount > 0 {
		nums[idx] = 2
		idx++
		blueCount--
	}
}
