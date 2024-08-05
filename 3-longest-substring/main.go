package main

import "fmt"

func main() {
	s := "dvdf"
	fmt.Println(lengthOfLongestSubstring(s))

	s = "bbbbb"
	fmt.Println(lengthOfLongestSubstring(s))

	// s = "pwwkew"
	// fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	} else if len(s) == 1 {
		return 1
	}

	var (
		tmp string = ""
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
