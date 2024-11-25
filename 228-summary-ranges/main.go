package main

import (
	"fmt"
	"strconv"
)

func main() {
	cases := []struct {
		id   int
		nums []int
		want []string
	}{
		{
			id:   1,
			nums: []int{0, 1, 2, 4, 5, 7},
			want: []string{"0->2", "4->5", "7"},
		},
		{
			id:   2,
			nums: []int{0, 2, 3, 4, 6, 8, 9},
			want: []string{"0", "2->4", "6", "8->9"},
		},
	}

	for _, c := range cases {
		res := summaryRanges(c.nums)
		if len(res) != len(c.want) {
			fmt.Printf("Test %d failed, len\n", c.id)
			continue
		}

		for i := range c.want {
			if res[i] != c.want[i] {
				fmt.Printf("Test %d failed\n", c.id)
				break
			}
		}
	}
	fmt.Println("Finish")
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	} else if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}

	res := []string{}
	start := 0
	prev := nums[start]

	i := 1

	for ; i < len(nums); i++ {
		prev = nums[i-1]
		if nums[i]-1 != prev {
			if i-start > 1 {
				res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[i-1]))
			} else {
				res = append(res, strconv.Itoa(nums[start]))
			}
			start = i
		}
	}

	if i-1 == start {
		res = append(res, strconv.Itoa(nums[start]))
	} else {
		res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[i-1]))
	}

	return res
}
