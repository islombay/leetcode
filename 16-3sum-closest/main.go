package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	cases := []struct {
		id     int
		nums   []int
		target int
		res    int
	}{
		{
			id:     1,
			nums:   []int{-1, 2, 1, -4},
			target: 1,
			res:    2,
		},
		{
			id:     2,
			nums:   []int{0, 0, 0},
			target: 1,
			res:    0,
		},
		{
			id:     3,
			nums:   []int{1, 1, 1, 0},
			target: 100,
			res:    3,
		},
		{
			id:     4,
			nums:   []int{4, 0, 5, -5, 3, 3, 0, -4, -5},
			target: -2,
			res:    -2,
		},
	}

	for _, c := range cases {
		if res := threeSumClosest(c.nums, c.target); res != c.res {
			fmt.Printf("Test %d failed, expected %d, got %d\n", c.id, c.res, res)
		}
	}
}
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closest := math.MaxInt32
	closestDiff := math.MaxInt32

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if abs(target-sum) < closestDiff {
				closest = sum
				closestDiff = abs(target - sum)
			}

			if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				return sum
			}
		}
	}
	return closest
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
