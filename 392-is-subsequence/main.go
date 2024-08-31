package main

import "fmt"

func main() {
	cases := []struct {
		name   string
		s      string
		t      string
		result bool
	}{
		{
			name:   "1",
			s:      "abc",
			t:      "ahbgdc",
			result: true,
		}, {
			name:   "2",
			s:      "axc",
			t:      "ahbgdc",
			result: false,
		}, {
			name:   "3",
			s:      "",
			t:      "asdflj",
			result: true,
		},
	}
	for _, c := range cases {
		if r := isSubsequence(c.s, c.t); r != c.result {
			fmt.Printf("- Test %s failed, expected %t, got %t\n", c.name, c.result, r)
		} else {
			fmt.Printf("Test %s passed, expected %t, got %t\n", c.name, c.result, r)
		}
	}
}

func isSubsequence(s string, t string) bool {
	if s == t {
		return true
	}

	if len(s) == 0 {
		return true
	}

	if len(s) > len(t) {
		return false
	}

	sPoint := 0

	for _, r := range t {
		if r == rune(s[sPoint]) {
			sPoint++
			if len(s) == sPoint {
				return true
			}
		}
	}

	return sPoint > len(s)
}
