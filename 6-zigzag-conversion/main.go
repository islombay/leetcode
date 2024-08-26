package main

import "fmt"

func main() {
	cases := []struct {
		name    string
		s       string
		numRows int
		result  string
	}{
		{
			name:    "1",
			s:       "PAYPALISHIRING",
			numRows: 3,
			result:  "PAHNAPLSIIGYIR",
		},
		{
			name:    "2",
			s:       "PAYPALISHIRING",
			numRows: 4,
			result:  "PINALSIGYAHRPI",
		},
		{
			name:    "3",
			s:       "A",
			numRows: 1,
			result:  "A",
		},
	}

	for _, c := range cases {
		if res := convert(c.s, c.numRows); res != c.result {
			fmt.Printf("- Test %s failed, expected %s, got %s\n", c.name, c.result, res)
		} else {
			fmt.Printf("Test %s passed, expected %s, got %s\n", c.name, c.result, res)
		}
	}
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	if numRows >= len(s) {
		return s
	}

	result := ""
	tmpIndex := 0

	for i := 0; i < numRows; i++ {
		tmpIndex = i
		for tmpIndex < len(s) {
			result += string(s[tmpIndex])

			if !(i == 0 || i == numRows-1) {
				tmp := tmpIndex + numRows - i + numRows - i - 2
				if tmp >= len(s) {
					break
				}

				result += string(s[tmp])
			}

			tmpIndex = tmpIndex + numRows + numRows - 2
		}
	}

	return result
}
