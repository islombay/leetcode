package main

import (
	"strconv"
	"testing"
)

func TestSearchMatrix(t *testing.T) {
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
		t.Run(strconv.Itoa(c.id), func(t *testing.T) {
			got := searchMatrix(c.matrix, c.target)
			if got != c.want {
				t.Errorf("got: %v, want: %v", got, c.want)
			}
		})
	}
}
