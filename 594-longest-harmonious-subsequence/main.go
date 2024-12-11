package main

import (
	"fmt"
	"sort"
)

func main() {
	cases := []struct {
		id   int
		nums []int
		want int
	}{
		{
			id:   1,
			nums: []int{1, 2, 2, 1},
			want: 4,
		},
		{
			id:   2,
			nums: []int{1, 1, 1, 1},
			want: 0,
		},
		{
			id:   3,
			nums: []int{1, 2, -1, 1, 2, 5, 2, 5, 2},
			want: 6,
		},
	}

	for _, c := range cases {
		if res := findLHS(c.nums); res != c.want {
			fmt.Printf("Test %d failed, expected %d, got %d\n", c.id, c.want, res)
		}
	}
}

func findLHS(nums []int) int {
	sort.Ints(nums)

	l, c := 0, 0
	prev := nums[l]
	for r := range nums {
		for l < r && nums[r] != prev && nums[r]-nums[l] != 1 {
			l++
			prev = nums[l]
		}
		if l < r && nums[r] != prev {
			c = max(c, r-l+1)

		}
	}
	return c
}
