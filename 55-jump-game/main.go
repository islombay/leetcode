package main

import "fmt"

func main() {
	cases := []struct {
		name   string
		nums   []int
		result bool
	}{
		{
			name:   "1",
			nums:   []int{2, 3, 1, 1, 4},
			result: true,
		},
		{
			name:   "2",
			nums:   []int{3, 2, 1, 0, 4},
			result: false,
		},
	}

	for _, c := range cases {
		result := canJump(c.nums)
		if result != c.result {
			fmt.Printf("Test %s failed, expected %t, got %t\n", c.name, c.result, result)
		} else {
			fmt.Printf("Test %s passed, expected %t, got %t\n", c.name, c.result, result)
		}
	}
}

type Index int

var (
	GOOD Index = 1
	BAD  Index = 2

	meta map[int]Index
)

func canJump(nums []int) bool {
	meta = make(map[int]Index)
	meta[len(nums)-1] = GOOD
	return canJumpFromPosition(0, nums)
}

func canJumpFromPosition(i int, nums []int) bool {
	if val, ok := meta[i]; ok && val == GOOD {
		return true
	} else if ok && val == BAD {
		return false
	}

	if nums[i] == 0 {
		meta[i] = BAD
		return false
	}

	for j := 1; j <= nums[i]; j++ {
		if i+nums[i] == len(nums)-1 {
			meta[i] = GOOD
			return true
		}
		if i+j < len(nums) {
			if canJumpFromPosition(i+j, nums) == true {
				meta[i] = GOOD
				return true
			}
		}
	}
	meta[i] = BAD
	return false
}
