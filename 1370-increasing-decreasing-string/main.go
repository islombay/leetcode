package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	cases := []struct {
		id   int
		s    string
		want string
	}{
		{
			id:   1,
			s:    "aaaabbbbcccc",
			want: "abccbaabccba",
		},
		{
			id:   1,
			s:    "rat",
			want: "art",
		},
	}

	for _, c := range cases {
		if res := sortString2(c.s); res != c.want {
			fmt.Printf("Test %d failed, expected %s, got %s\n", c.id, c.want, res)
		}
	}
}

type RuneCounter struct {
	Char  rune
	Count int
}

func sortString2(s string) string {
	charCounter := make([]RuneCounter, 26)
	for _, c := range s {
		idx := int(c) - int('a')
		charCounter[idx].Char = c
		charCounter[idx].Count++
	}

	// aaabbbccc
	// a: 4, b: 4, c: 4,
	res := strings.Builder{}
	// abccba

	for len(charCounter) > 0 {
		for i := 0; i < len(charCounter); i++ {
			if charCounter[i].Count > 0 {
				res.WriteRune(charCounter[i].Char)
				charCounter[i].Count--
			} else if charCounter[i].Count == 0 {
				charCounter = append(charCounter[:i], charCounter[i+1:]...)
				i--
			}
		}

		for i := len(charCounter) - 1; i >= 0; i-- {
			if charCounter[i].Count > 0 {
				res.WriteRune(charCounter[i].Char)
				charCounter[i].Count--
			} else if charCounter[i].Count == 0 {
				charCounter = append(charCounter[:i], charCounter[i+1:]...)
			}
		}
	}

	return res.String()
}

func sortString(s string) string {
	res := ""

	if len(s) == 1 {
		return s
	}

	runes := []rune(s)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	s = string(runes)

	for s != "" {
		prev := s[0]
		s = s[1:]
		res += string(prev)

		for i := 0; i < len(s); i++ {
			if s[i] > prev {
				res += string(s[i])
				prev = s[i]
				s = s[0:i] + s[i+1:]
				i--
			}
		}

		if len(s) == 0 {
			return res
		}

		prev = s[len(s)-1]
		s = s[:len(s)-1]
		res += string(prev)

		for i := len(s) - 1; i >= 0; i-- {
			if s[i] < prev {
				res += string(s[i])
				prev = s[i]
				s = s[0:i] + s[i+1:]
			}
		}
	}

	return res
}
