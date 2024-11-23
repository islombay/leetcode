package main

import "fmt"

func main() {
	cases := []struct {
		name   string
		s, t   string
		result bool
	}{
		{
			name:   "1",
			s:      "anagram",
			t:      "nagaram",
			result: true,
		},
		{
			name:   "2",
			s:      "tar",
			t:      "car",
			result: false,
		},
	}

	for _, c := range cases {
		if res := isAnagram(c.s, c.t); res != c.result {
			fmt.Printf("- Test %s failed, expected %t, got %t\n", c.name, c.result, res)
		} else {
			fmt.Printf("Test %s passed, expected %t, got %t\n", c.name, c.result, res)
		}
	}
}

func isAnagram(s string, t string) bool {
	var (
		list = make(map[rune]int)
	)

	for _, r := range s {
		list[r]++
	}

	for _, r := range t {
		list[r]--
		if list[r] < 0 {
			return false
		}

		if val, ok := list[r]; val == 0 && ok {
			delete(list, r)
		}
	}

	if len(list) != 0 {
		return false
	}

	return true
}
