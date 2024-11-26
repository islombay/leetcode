package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		id   int
		s    string
		want int
	}{
		{
			id:   1,
			s:    "leetcode",
			want: 0,
		},
		{
			id:   2,
			s:    "loveleetcode",
			want: 2,
		},
		{
			id:   3,
			s:    "aabb",
			want: -1,
		},
	}

	for _, c := range cases {
		res := firstUniqChar(c.s)
		if res != c.want {
			fmt.Printf("Test %d failed, expected %d, got %d\n", c.id, c.want, res)
		}
	}
	fmt.Println("Finish")
}

func firstUniqChar(s string) int {
	m := make(map[rune]int)
	for _, c := range s {
		m[c]++
	}

	for i, c := range s {
		if m[c] == 1 {
			return i
		}
	}
	return -1
}
