package main

import "fmt"

func main() {
	cases := []struct {
		name   string
		s      string
		result int
	}{
		{
			name:   "1",
			s:      "abcabcbb",
			result: 3,
		},
		{
			name:   "2",
			s:      "bbbbb",
			result: 1,
		},
		{
			name:   "3",
			s:      "pwwkew",
			result: 3,
		},
	}

	for _, c := range cases {
		if res := lengthOfLongestSubstring_2(c.s); res != c.result {
			fmt.Printf("- Test %s failed, expected %d, got %d\n", c.name, c.result, res)
		} else {
			fmt.Printf("Test %s passed, expected %d, got %d\n", c.name, c.result, res)
		}
	}
}

func lengthOfLongestSubstring_2(s string) int {
	var (
		//unq       = make(map[uint8]int)
		start     = 0
		end       = 0
		maxLength = 0
		//tmp       = ""
	)

	for ; end < len(s); end++ {
		if validSTR(s[start : end+1]) {
			maxLength = max(maxLength, end-start+1)
		} else {
			start++

		}
	}

	return maxLength
}

func validSTR(s string) bool {
	m := make(map[rune]int)
	for _, c := range s {
		if _, ok := m[c]; ok {
			return false
		} else {
			m[c] = 1
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	} else if len(s) == 1 {
		return 1
	}

	var (
		tmp = ""
	)

	subs := make(map[string]int, 0)

Outer:
	for i := 0; i < len(s); i++ {
		// fmt.Println()
		// fmt.Println("Current I value:", string(s[i]), "- subs:", subs)
		for j := i; j < len(s); j++ {
			// fmt.Println("Current J value:", string(s[j]), "- tmp:", tmp)
			for ind := range tmp {
				if tmp[ind] == s[j] {
					subs[tmp] = len(tmp)
					tmp = ""
					continue Outer
				}
			}

			tmp += string(s[j])
		}
	}

	if tmp != "" {
		subs[tmp] = len(tmp)
	}

	var (
		maxValue int
		maxKey   string
		first    bool = true
	)

	for key, val := range subs {
		if first || val > maxValue {
			maxValue = val
			maxKey = key
			first = false
		}
	}
	return subs[maxKey]
}
