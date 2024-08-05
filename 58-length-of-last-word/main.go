package main

import "fmt"

func main() {
	str := "ada "

	fmt.Println(lengthOfLastWord(str))
}

func lengthOfLastWord(s string) int {
	var space_rune rune = 32

	if len(s) == 1 {
		return 1
	}

	var (
		endIndex int = -1
	)

	for i := len(s) - 1; i > -1; i-- {
		if s[i] == byte(space_rune) {
			if endIndex != -1 {
				fmt.Println("#a here")
				return endIndex - i
			}
		} else if endIndex == -1 {
			endIndex = i
		}
	}

	fmt.Println("#b here")
	if endIndex == 0 {
		return 1
	}
	return endIndex + 1
}
