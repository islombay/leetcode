package main

import (
	"fmt"
	"reflect"
)

func main() {
	cases := []struct {
		nums []int
		res  []int
		id   int
	}{
		{
			id:   1,
			nums: []int{0, 2, 1, 5, 3, 4},
			res:  []int{0, 1, 2, 4, 5, 3},
		},
		{
			id:   2,
			nums: []int{3, 2, 1, 0},
			res:  []int{0, 1, 2, 3},
		},
	}

	for _, c := range cases {
		r := buildArray(c.nums)
		if !reflect.DeepEqual(r, c.res) {
			fmt.Printf("Fail %d, want: %v, got %v\n", c.id, c.res, r)
		}
	}
}

func buildArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i, _ := range nums {
		ans[i] = nums[nums[i]]
	}
	return ans
}
