package main

import "fmt"

func main() {
	cases := []struct {
		id     int
		digits string
		res    []string
	}{
		{
			id:     1,
			digits: "23",
			res:    []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			id:     2,
			digits: "",
			res:    []string{},
		},
		{
			id:     3,
			digits: "2",
			res:    []string{"a", "b", "c"},
		},
		{
			id:     4,
			digits: "22",
			res:    []string{"aa", "ab", "ac", "ba", "bb", "bc", "ca", "cb", "cc"},
		},
	}

	for _, c := range cases {
		res := letterCombinations(c.digits)
		if len(res) != len(c.res) {
			fmt.Printf("Test %d failed, len\n", c.id)
			continue
		}

		for i := 0; i < len(res); i++ {
			if res[i] != c.res[i] {
				fmt.Printf("Test %d failed\n", c.id)
				break
			}
		}
	}
}

func letterCombinations(digits string) []string {
	letters := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	res := []string{}

	// 2
	for i, r := range digits {
		lts := letters[r]
		tmp := ""
		for _, letter := range lts {
			if len(digits) == 1 {
				tmp += string(letter)
			}
			for _, rn := range digits[i+1:] {
				secondaryLetters := letters[rn]
				for _, letter2 := range secondaryLetters {
					if tmp == "" {
						tmp += string(letter)
					}
					tmp += string(letter2)
					res = append(res, tmp)
					tmp = ""
				}
			}
			if tmp != "" {
				res = append(res, tmp)
				tmp = ""
			}
		}
	}
	return res
}
