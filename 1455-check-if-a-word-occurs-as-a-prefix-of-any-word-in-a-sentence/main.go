package main

import (
	"fmt"
	"strings"
)

func main() {
	cases := []struct {
		id                   int
		sentence, searchWord string
		want                 int
	}{
		{
			id:         1,
			sentence:   "i love eating burger",
			searchWord: "burg",
			want:       4,
		},
		{
			id:         2,
			sentence:   "this problem is an easy problem",
			searchWord: "pro",
			want:       2,
		},
		{
			id:         3,
			sentence:   "i am tired",
			searchWord: "you",
			want:       -1,
		},
		{
			id:         4,
			sentence:   "hellohello hellohellohello",
			searchWord: "ell",
			want:       -1,
		},
		{
			id:         5,
			sentence:   "i am your dad",
			searchWord: "daddy",
			want:       -1,
		},
		{
			id:         6,
			sentence:   "b bu bur burg burger",
			searchWord: "burg",
			want:       4,
		},
	}

	for _, c := range cases {
		if res := isPrefixOfWord2(c.sentence, c.searchWord); res != c.want {
			fmt.Printf("Test %d failed, expected %d, got %d\n", c.id, c.want, res)
		}
	}
}

func isPrefixOfWord2(sentence string, searchWord string) int {
	// do not split, and check as it is (save memory)
	if len(sentence) < 2 {
		if sentence == searchWord {
			return 1
		}
		return -1
	}
	l, r := 0, 1
	spaceCount := 0

	for l < len(sentence) && r < len(sentence) {
		for r < len(sentence) && sentence[r] != ' ' {
			r++
		}
		spaceCount++
		if strings.HasPrefix(sentence[l:r], searchWord) {
			return spaceCount
		}
		l = r + 1
		r += 2
	}
	return -1
}

func isPrefixOfWord(sentence string, searchWord string) int {
	//Solution where i split and then check (memory consuming)
	strs := strings.Split(sentence, " ")
	for i, str := range strs {
		if len(str) < len(searchWord) {
			continue
		}
		failed := false
		for j := range searchWord {
			if str[j] != searchWord[j] {
				failed = true
				break
			}
		}
		if !failed {
			return i + 1
		}
	}
	return -1
}
