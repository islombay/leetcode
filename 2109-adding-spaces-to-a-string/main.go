package main

import (
	"fmt"
	"strings"
)

func main() {
	cases := []struct {
		id     int
		s      string
		spaces []int
		want   string
	}{
		{
			id:     1,
			s:      "LeetcodeHelpsMeLearn",
			spaces: []int{8, 13, 15},
			want:   "Leetcode Helps Me Learn",
		},
		{
			id:     2,
			s:      "icodeinpython",
			spaces: []int{1, 5, 7, 9},
			want:   "i code in py thon",
		},
		{
			id:     3,
			s:      "spacing",
			spaces: []int{0, 1, 2, 3, 4, 5, 6},
			want:   " s p a c i n g",
		},
	}

	for _, c := range cases {
		if res := addSpaces(c.s, c.spaces); res != c.want {
			fmt.Printf("Test %d failed, expected: (%s), got (%s)\n", c.id, c.want, res)
		}
	}
}

func addSpaces(s string, spaces []int) string {
	newString := strings.Builder{}
	lastPos := 0
	for _, v := range spaces {
		newString.WriteString(s[lastPos:v])
		newString.WriteRune(' ')
		lastPos = v
	}
	if lastPos < len(s) {
		newString.WriteString(s[lastPos:])
	}
	return newString.String()
}
