package main

import (
	"fmt"
	"sort"
)

func main() {
	cases := []struct {
		id   int
		nums []int
		want bool
	}{
		{
			id:   1,
			nums: []int{1, 4, 3, 1},
			want: true,
		},
		{
			id:   2,
			nums: []int{1, 2, 3, 4},
			want: false,
		},
		{
			id:   3,
			nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			want: true,
		},
	}

	for _, c := range cases {
		if containsDuplicate(c.nums) != c.want {
			fmt.Printf("Test %d failed\n", c.id)
		}
	}
	fmt.Println("Finished")
}

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	sort.Ints(nums)
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == prev {
			return true
		}
		prev = nums[i]
	}
	return false
}
