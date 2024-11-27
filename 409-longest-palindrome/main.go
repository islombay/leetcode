package main

import "fmt"

func main() {
	cases := []struct {
		id   int
		s    string
		want int
	}{
		{
			id:   1,
			s:    "abccccdd",
			want: 7,
		},
		{
			id:   2,
			s:    "a",
			want: 1,
		},
		{
			id:   3,
			s:    "Aa",
			want: 1,
		},
		{
			id:   4,
			s:    "Aabb",
			want: 3,
		},
		{
			id:   5,
			s:    "aaabbc",
			want: 5,
		},
	}
	for _, c := range cases {
		if res := longestPalindrome(c.s); res != c.want {
			fmt.Printf("Test %d failed, expected %d, got %d\n", c.id, c.want, res)
		}
	}
	fmt.Println("Finished")
}

func longestPalindrome(s string) int {
	m := map[rune]int{}
	count := 0
	singleElementUsed := false

	for _, c := range s {
		m[c]++
	}
	for _, v := range m {
		if v == 1 && singleElementUsed {
			continue
		}
		count += v - (v & 1)

		if v&1 == 1 {
			if !singleElementUsed {
				count++
				singleElementUsed = true
			}
		}
	}
	return count
}
