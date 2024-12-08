package main

import (
	"fmt"
	"sort"
)

func main() {
	cases := []struct {
		id     int
		events [][]int
		want   int
	}{
		{
			id:     1,
			events: [][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}},
			want:   4,
		},
		{
			id:     2,
			events: [][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}},
			want:   5,
		},
		{
			id:     3,
			events: [][]int{{1, 5, 3}, {1, 5, 1}, {6, 6, 5}},
			want:   8,
		},
		{
			id:     4,
			events: [][]int{{35, 90, 47}, {72, 80, 70}},
			want:   70,
		},
	}
	for _, c := range cases {
		if res := maxTwoEvents(c.events); res != c.want {
			fmt.Printf("Test %d failed, expected %d, got %d\n", c.id, c.want, res)
		}
	}
}

func maxTwoEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})

	maxValues := make([]int, len(events))
	maxValues[0] = events[0][2]

	maxSum := 0

	for i := 1; i < len(events); i++ {
		maxValues[i] = max(maxValues[i-1], events[i][2])
	}

	for _, event := range events {
		if idx := search(events, event[0]-1); idx != -1 {
			event[2] += maxValues[idx]
		}
		maxSum = max(maxSum, event[2])
	}
	return maxSum
}

func search(events [][]int, target int) int {
	l, r, res := 0, len(events)-1, -1

	for l <= r {
		mid := l + (r-l)/2
		if events[mid][1] <= target {
			res = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return res
}
