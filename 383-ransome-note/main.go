package main

import "fmt"

func main() {
	cases := []struct {
		name       string
		ransomNote string
		magazine   string
		result     bool
	}{
		{
			name:       "1",
			ransomNote: "a",
			magazine:   "b",
			result:     false,
		},
		{
			name:       "2",
			ransomNote: "aa",
			magazine:   "ab",
			result:     false,
		},
		{
			name:       "3",
			ransomNote: "aa",
			magazine:   "aab",
			result:     true,
		},
	}

	for _, c := range cases {
		if result := canConstruct(c.ransomNote, c.magazine); result != c.result {
			fmt.Printf("- Test %s failed, expected %t, got %t\n", c.name, c.result, result)
		} else {
			fmt.Printf("Test %s passed, expected %t, got %t\n", c.name, c.result, result)
		}
	}
}

func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}

	letters := make(map[int32]int)
	for _, letter := range magazine {
		letters[letter]++
	}

	for _, letter := range ransomNote {
		val, ok := letters[letter]
		if !ok || val <= 0 {
			return false
		}
		if val > 0 {
			letters[letter]--
		}
	}

	return true
}
