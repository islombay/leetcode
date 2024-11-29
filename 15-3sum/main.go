package main

import (
	"fmt"
	"sort"
)

func main() {
	cases := []struct {
		id   int
		nums []int
		res  [][]int
	}{
		{
			id:   1,
			nums: []int{-1, 0, 1, 2, -1, -4},
			res: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			id:   2,
			nums: []int{0, 1, 1},
			res:  [][]int{},
		},
		{
			id:   3,
			nums: []int{0, 0, 0},
			res:  [][]int{{0, 0, 0}},
		},
	}

	for _, c := range cases {
		res := threeSum(c.nums)
		if len(res) != len(c.res) {
			fmt.Printf("Test %d failed, len\n", c.id)
			continue
		}
		failed := false
		for i := 0; i < len(c.res) && !failed; i++ {
			if len(c.res[i]) != len(res[i]) {
				fmt.Printf("Test %d failed, len\n", c.id)
				break
			}
			for j := 0; j < len(c.res[i]); j++ {
				if c.res[i][j] != res[i][j] {
					fmt.Printf("Test %d failed\n", c.id)
					failed = true
					break
				}
			}
		}
	}
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return res
}
