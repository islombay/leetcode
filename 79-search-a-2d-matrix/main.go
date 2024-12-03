package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		id     int
		matrix [][]int
		target int
		want   bool
	}{
		{
			id:     1,
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 3,
			want:   true,
		},
		{
			id:     2,
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 13,
			want:   false,
		},
		{
			id:     3,
			matrix: [][]int{{1}, {3}},
			target: 3,
			want:   true,
		},
	}

	for _, c := range cases {
		got := searchMatrix(c.matrix, c.target)
		if got != c.want {
			fmt.Printf("Test %d failed, expected %t, got %t\n", c.id, c.want, got)
		}
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	subLen := len(matrix[0])
	l, r := 0, len(matrix)*subLen-1
	for l <= r {
		mid := l + (r-l)/2
		if matrix[mid/subLen][mid%subLen] == target {
			return true
		} else if matrix[mid/subLen][mid%subLen] < target {
			l = mid + 1
		} else if matrix[mid/subLen][mid%subLen] > target {
			r = mid - 1
		}
	}
	return false
}
