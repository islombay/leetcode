package main

import (
	"fmt"
	"sort"
)

func main() {
	cases := []struct {
		id   int
		g    []int
		s    []int
		want int
	}{
		{
			id:   1,
			g:    []int{1, 2, 3},
			s:    []int{1, 1},
			want: 1,
		},
		{
			id:   2,
			g:    []int{1, 2},
			s:    []int{1, 2, 3},
			want: 2,
		},
		{
			id:   3,
			g:    []int{10, 9, 8, 7},
			s:    []int{5, 6, 7, 8},
			want: 2,
		},
	}

	for _, c := range cases {
		if res := findContentChildren(c.g, c.s); res != c.want {
			fmt.Printf("Test %d failed, expected %d, got %d\n", c.id, c.want, res)
		}
	}
	fmt.Println("Finish")
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	count := 0
	gIdx := 0
	sIdx := 0

	// 7 8 9 10
	// 5 6 7 8
	for gIdx < len(g) && sIdx < len(s) {
		if s[sIdx] >= g[gIdx] {
			count++
			sIdx++
			gIdx++
			continue
		}
		sIdx++
	}
	return count
}
