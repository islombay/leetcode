package main

import "fmt"

func main() {
	cases := []struct {
		id         int
		timeSeries []int
		duration   int
		want       int
	}{
		{
			id:         1,
			timeSeries: []int{1, 4},
			duration:   2,
			want:       4,
		},
		{
			id:         2,
			timeSeries: []int{1, 2},
			duration:   2,
			want:       3,
		},
		{
			id:         3,
			timeSeries: []int{1, 2, 9},
			duration:   5,
			want:       11,
		},
	}
	for _, c := range cases {
		if res := findPoisonedDuration(c.timeSeries, c.duration); res != c.want {
			fmt.Printf("Test %d failed, expected %d, got %d\n", c.id, c.want, res)
		}
	}
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	total := 0
	timerLast := 0
	for _, t := range timeSeries {
		total += duration
		if t <= timerLast {
			total -= timerLast - t + 1
		}
		timerLast = t + duration - 1
	}
	return total
}
