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
			nums: []int{0, 1},
			want: 2,
		},
		{
			id:   2,
			nums: []int{3, 0, 1},
			want: 2,
		},
		{
			id:   3,
			nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			want: 8,
		},
	}
	for _, c := range cases {
		res := missingNumber(c.nums)
		if res != c.want {
			fmt.Printf("Test %d failed, expected %d, got %d\n", c.id, c.want, res)
		}
	}
	fmt.Println("Finished")
}

func missingNumber(nums []int) int {
	sort.Ints(nums)
	if len(nums) == 1 {
		if nums[0] != 0 {
			return 0
		}
		return 1
	}

	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]-1 != prev {
			return nums[i] - 1
		}
		prev = nums[i]
	}

	if nums[len(nums)-1] != len(nums) {
		return len(nums)
	}
	return 0
}
