package main

import (
	"fmt"
	"reflect"
)

func main() {
	cases := []struct {
		id    int
		nums  []int
		index []int
		want  []int
	}{
		{
			id:    1,
			nums:  []int{4, 2, 4, 3, 2},
			index: []int{0, 0, 1, 3, 1},
			want:  []int{2, 2, 4, 4, 3},
		},
	}

	for _, c := range cases {
		res := createTargetArray(c.nums, c.index)
		if !reflect.DeepEqual(res, c.want) {
			fmt.Printf("Test %d failed\n", c.id)
		}
	}
}

func createTargetArray(nums []int, index []int) []int {
	res := make([]int, len(nums))

	for idx, i := range index {
		for j := len(res) - 1; j > i; j-- {
			res[j] = res[j-1]
		}

		res[i] = nums[idx]
	}

	return res
}
