package main

import "fmt"

func main() {
	cases := []struct {
		name   string
		s      string
		output int
	}{
		{
			name:   "1",
			s:      "III",
			output: 3,
		},
		{
			name:   "2",
			s:      "LVIII",
			output: 58,
		},
		{
			name:   "3",
			s:      "MCMXCIV",
			output: 1994,
		},
	}

	for _, c := range cases {
		result := romanToInt(c.s)
		if result != c.output {
			fmt.Printf("- Test %s failed, expected %d, got %d\n", c.name, c.output, result)
		} else {
			fmt.Printf("Test %s passed, expected %d, got %d\n", c.name, c.output, result)
		}
	}
}

func romanToInt(s string) int {
	nums := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var (
		tmp    int
		result int
	)

	for i := len(s) - 1; i >= 0; i-- {
		val, _ := nums[s[i]]
		if val >= tmp {
			tmp = val
			result += val
		} else {
			result -= val
		}
	}

	return result
}
