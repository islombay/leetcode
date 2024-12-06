package main

import "fmt"

func main() {
	cases := []struct {
		id     int
		banned []int
		n      int
		maxSum int
		want   int
	}{
		{
			id:     1,
			banned: []int{1, 6, 5},
			n:      5,
			maxSum: 6,
			want:   2,
		},
		{
			id:     2,
			banned: []int{1, 2, 3, 4, 5, 6, 7},
			n:      8,
			maxSum: 1,
			want:   0,
		},
		{
			id:     3,
			banned: []int{11},
			n:      7,
			maxSum: 50,
			want:   7,
		},
	}

	for _, c := range cases {
		if res := maxCount(c.banned, c.n, c.maxSum); res != c.want {
			fmt.Printf("Test %d failed, expected %d, got %d\n", c.id, c.want, res)
		}
	}
}

func maxCount(banned []int, n int, maxSum int) int {
	m := map[int]bool{}
	for _, v := range banned {
		if v > n {
			continue
		}
		m[v] = true
	}

	totalSum := 0
	total := 0

	for i := 1; i <= n; i++ {
		if _, ok := m[i]; ok {
			continue
		}
		totalSum += i
		if totalSum > maxSum {
			return total
		}
		total++
	}

	return total
}
