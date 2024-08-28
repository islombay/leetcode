package main

import (
	"fmt"
	"strings"
)

func main() {
	cases := []struct {
		name    string
		pattern string
		s       string
		result  bool
	}{
		{
			name:    "1",
			pattern: "abba",
			s:       "dog cat cat dog",
			result:  true,
		},
		{
			name:    "2",
			pattern: "abba",
			s:       "dog cat cat fish",
			result:  false,
		},
		{
			name:    "3",
			pattern: "aaaa",
			s:       "dog cat cat dog",
			result:  false,
		},
		{
			name:    "4",
			pattern: "abba",
			s:       "dog dog dog dog",
			result:  false,
		},
		{
			name:    "5",
			pattern: "abc",
			s:       "dog cat dog",
			result:  false,
		},
	}

	for _, c := range cases {
		if result := wordPattern(c.pattern, c.s); result != c.result {
			fmt.Printf("- Test %s failed, expected %t, got %t\n", c.name, c.result, result)
		} else {
			fmt.Printf("Test %s passed, expected %t, got %t\n", c.name, c.result, result)
		}
	}
}

func wordPattern(pattern string, s string) bool {
	words := make(map[byte]string)
	reverse := make(map[string]byte)

	sArray := strings.Split(s, " ")

	if len(sArray) < len(pattern) || len(sArray) > len(pattern) {
		return false
	}

	for i := 0; i < len(sArray); i++ {
		val, ok := words[pattern[i]]
		reverseVal, reverseOk := reverse[sArray[i]]

		if !ok && reverseOk {
			return false
		}
		if !ok {
			words[pattern[i]] = sArray[i]
			reverse[sArray[i]] = pattern[i]
			continue
		}
		if !ok {
			return false
		}

		if val != sArray[i] || reverseVal != pattern[i] {
			return false
		}
	}
	return true
}
