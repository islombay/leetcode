package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		name string
		s    string
		n    int
	}{
		{
			name: "1",
			s:    "Hello World",
			n:    5,
		},
		{
			name: "2",
			s:    "   fly me   to   the moon  ",
			n:    4,
		},
		{
			name: "3",
			s:    "luffy is still joyboy",
			n:    6,
		},
	}

	for _, c := range cases {
		r := lengthOfLastWord(c.s)
		if r != c.n {
			fmt.Printf("- Test %s failed, expected %d, got %d\n", c.name, c.n, r)
		} else {
			fmt.Printf("Test %s passed, expected %d, got %d\n", c.name, c.n, r)
		}
	}
}

//func lengthOfLastWord(s string) int {
//	s = strings.TrimSpace(s)
//	a := strings.Split(s, " ")
//	return len(a[len(a)-1])
//}

func lengthOfLastWord(s string) int {
	var space_rune rune = 32

	if len(s) == 1 {
		return 1
	}

	var (
		endIndex int = -1
	)

	for i := len(s) - 1; i > -1; i-- {
		if s[i] == byte(space_rune) {
			if endIndex != -1 {
				fmt.Println("#a here")
				return endIndex - i
			}
		} else if endIndex == -1 {
			endIndex = i
		}
	}

	fmt.Println("#b here")
	if endIndex == 0 {
		return 1
	}
	return endIndex + 1
}
