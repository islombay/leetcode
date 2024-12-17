package main

func sortedSquares(nums []int) []int {
	minIdx := 0
	squares := make([]int, len(nums))

	for i, n := range nums {
		if abs(nums[minIdx]) > abs(n) {
			minIdx = i
		}

		squares[i] = n * n
	}

	l, r := minIdx, minIdx+1
	i := 0
	for (l >= 0 || r < len(nums)) && i < len(nums) {
		if l >= 0 && r < len(nums) {
			if squares[l] > squares[r] {
				nums[i] = squares[r]
				r++
			} else {
				nums[i] = squares[l]
				l--
			}
		} else if l >= 0 {
			nums[i] = squares[l]
			l--
		} else {
			nums[i] = squares[r]
			r++
		}
		i++
	}

	return nums
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
