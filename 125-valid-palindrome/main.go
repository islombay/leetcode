// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	cases := []struct {
		name   string
		s      string
		result bool
	}{
		{
			name:   "1",
			s:      "A man, a plan, a canal: Panama",
			result: true,
		}, {
			name:   "2",
			s:      "race a car",
			result: false,
		}, {
			name:   "3",
			s:      " ",
			result: true,
		}, {
			name:   "4",
			s:      "0P",
			result: false,
		},
	}

	for _, c := range cases {
		if r := isPalindrome(c.s); r != c.result {
			fmt.Printf("- Test %s failed, expected %t, got %t\n", c.name, c.result, r)
		} else {
			fmt.Printf("Test %s passed, expected %t, got %t\n", c.name, c.result, r)
		}
	}
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	if len(s) == 0 {
		return true
	}

	var result string
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result += string(r)
		}
	}

	left, right := 0, len(result)-1

	for left <= right {
		if result[left] != result[right] {
			return false
		}
		left++
		right--
	}

	return true
}
