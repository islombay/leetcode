package main

import "fmt"

func main() {
	cases := []struct {
		name   string
		s      string
		result string
	}{
		{
			name:   "1",
			s:      "the sky is blue",
			result: "blue is sky the",
		},
		{
			name:   "2",
			s:      "   hello world   ",
			result: "world hello",
		},
		{
			name:   "3",
			s:      "a good   example",
			result: "example good a",
		},
		{
			name:   "4",
			s:      "EPY2giL",
			result: "EPY2giL",
		},
	}

	for _, c := range cases {
		if result := reverseWords(c.s); result != c.result {
			fmt.Printf("- Test %s failed, expected %s, got %s\n", c.name, c.result, result)
		} else {
			fmt.Printf("Test %s passed, expected %s, got %s\n", c.name, c.result, result)
		}
	}
}

func reverseWords(s string) string {
	n := len(s)
	tmp := ""
	result := ""

	for i := n - 1; i >= 0; i-- {
		if s[i] != ' ' {
			tmp = string(s[i]) + tmp
		} else if s[i] == ' ' && len(tmp) != 0 {
			if len(result) == 0 {
				result += tmp
			} else {
				result += " " + tmp
			}
			tmp = ""
		}
	}

	if tmp != "" {
		if len(result) == 0 {
			result += tmp
		} else {
			result += " " + tmp
		}
	}

	return result
}
