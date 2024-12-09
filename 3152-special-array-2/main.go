package main

import (
	"fmt"
	"reflect"
)

func main() {
	cases := []struct {
		id      int
		nums    []int
		queries [][]int
		want    []bool
	}{
		{
			id:      1,
			nums:    []int{4, 3, 1, 6},
			queries: [][]int{{0, 2}, {2, 3}},
			want:    []bool{false, true},
		},
		{
			id:      2,
			nums:    []int{3, 4, 2, 7, 4, 5, 6, 9, 10, 12, 13},
			queries: [][]int{{0, 5}, {9, 10}, {4, 8}, {1, 4}},
			want:    []bool{false},
		},
		{
			id:      3,
			nums:    []int{2, 8, 10, 9},
			queries: [][]int{{1, 3}},
			want:    []bool{false},
		},
	}

	for _, c := range cases {
		res := isArraySpecial(c.nums, c.queries)
		if !reflect.DeepEqual(res, c.want) {
			fmt.Printf("Test %d failed\n", c.id)
		}
	}
}

func isArraySpecial(nums []int, queries [][]int) []bool {
	answer := make([]bool, len(queries))

	badIdxs := []int{}
	prevEven := nums[0]&1 == 0
	for i := 1; i < len(nums); i++ {
		thisEven := nums[i]&1 == 0
		if thisEven == prevEven {
			badIdxs = append(badIdxs, i)
		}
		prevEven = thisEven
	}

	for idx, query := range queries {
		answer[idx] = !isBad(badIdxs, query)
	}

	return answer
}

func isBad(badIdxs []int, event []int) bool {
	l, r := 0, len(badIdxs)-1

	for l <= r {
		mid := l + (r-l)/2
		if event[0] < badIdxs[mid] && badIdxs[mid] <= event[1] {
			return true
		}
		if badIdxs[mid] <= event[0] {
			// r = mid - 1
			l = mid + 1
		} else {
			// l = mid + 1
			r = mid - 1
		}
	}
	return false
}
