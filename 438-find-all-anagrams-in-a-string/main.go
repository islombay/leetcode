package main

import (
	"fmt"
	"sort"
)

func main() {
	cases := []struct {
		id   int
		s, p string
		want []int
	}{
		{
			id:   1,
			s:    "cbaebabacd",
			p:    "abc",
			want: []int{0, 6},
		}, {
			id:   2,
			s:    "abab",
			p:    "ab",
			want: []int{0, 1, 2},
		},
	}

	for _, c := range cases {
		res := findAnagrams(c.s, c.p)
		if len(res) != len(c.want) {
			fmt.Printf("Test %d failed, len\n", c.id)
		}
		sort.Ints(res)
		sort.Ints(c.want)
		for i, v := range res {
			if v != c.want[i] {
				fmt.Printf("Test %d failed\n", c.id)
				break
			}
		}
	}
}

type Key [26]byte

func makeHash(s string) Key {
	k := Key{}
	for i := 0; i < len(s); i++ {
		k[s[i]-'a']++
	}
	return k
}

func findAnagrams(s string, p string) []int {
	hash := makeHash(p)

	res := []int{}

	for i := 0; i <= len(s)-len(p); i++ {
		newHash := makeHash(s[i : i+len(p)])
		if newHash == hash {
			res = append(res, i)
		}
	}
	return res
}
