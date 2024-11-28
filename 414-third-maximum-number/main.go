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
			nums: []int{1, 2, 3},
			want: 1,
		},
		{
			id:   2,
			nums: []int{2, 2, 3, 1},
			want: 1,
		},
		{
			id:   3,
			nums: []int{3, 1},
			want: 3,
		},
	}

	for _, c := range cases {
		if res := thirdMax(c.nums); res != c.want {
			fmt.Printf("Test %d failed, expected %d, got %d\n", c.id, c.want, res)
		}
	}
	fmt.Println("Finish")
}

func thirdMax(nums []int) int {
	sort.Ints(nums)

	if len(nums) < 3 {
		return nums[len(nums)-1]
	}

	prev := nums[len(nums)-1]
	counted := 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] == prev {
			continue
		}
		prev = nums[i]
		counted++
		if counted == 3 {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}
