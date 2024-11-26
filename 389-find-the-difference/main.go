package main

import "fmt"

func main() {
	cases := []struct {
		id   int
		s, t string
		want byte
	}{
		{
			id:   1,
			s:    "abcd",
			t:    "abcde",
			want: 'e',
		},
		{
			id:   2,
			s:    "",
			t:    "y",
			want: 'y',
		},
		{
			id:   3,
			s:    "a",
			t:    "aa",
			want: 'a',
		},
	}
	for _, c := range cases {
		if res := findTheDifference(c.s, c.t); res != c.want {
			fmt.Printf("Test %d failed, expected %s, got %s\n", c.id, string(c.want), string(res))
		}
	}
	fmt.Println("Finished")
}

func findTheDifference(s string, t string) byte {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}

	for _, v := range t {
		if m[v] == 0 {
			return byte(v)
		}
		m[v]--
	}
	return '0'
}
