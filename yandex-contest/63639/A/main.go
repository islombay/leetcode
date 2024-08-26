package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	line = strings.TrimSpace(line)

	wrapper := []string{}
	maxWordLen := 0
	tmpWord := ""

	for _, r := range line {
		if unicode.IsLetter(r) {
			tmpWord += string(r)
		} else if r == ',' || r == ' ' {
			if len(tmpWord) == 0 && r == ',' {
				wrapper[len(wrapper)-1] = wrapper[len(wrapper)-1] + string(r)
			} else if len(tmpWord) != 0 {
				if maxWordLen < len(tmpWord) {
					maxWordLen = len(tmpWord)
				}
				if r == ',' {
					tmpWord += string(r)
				}
				wrapper = append(wrapper, tmpWord)
				tmpWord = ""
			}
		}
	}
	if tmpWord != "" {
		if maxWordLen < len(tmpWord) {
			maxWordLen = len(tmpWord)
		}
		wrapper = append(wrapper, tmpWord)
	}

	maxLineLen := maxWordLen * 3
	tmpWord = ""
	resLines := []string{}

	for _, word := range wrapper {
		tmpLen := len(tmpWord) + len(word)
		if tmpWord != "" {
			tmpLen += 1
		}

		if tmpLen > maxLineLen {
			resLines = append(resLines, tmpWord)
			tmpWord = word
		} else {
			if tmpWord != "" {
				tmpWord += " " + word
			} else {
				tmpWord = word
			}
		}
	}
	if tmpWord != "" {
		resLines = append(resLines, tmpWord)
	}

	for _, line := range resLines {
		fmt.Println(line)
	}
}
