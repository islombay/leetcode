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
			s:    "aaaa",
			want: 2,
		},
		{
			id:   2,
			s:    "abcdef",
			want: -1,
		},
		{
			id:   3,
			s:    "abcaba",
			want: 1,
		},
		{
			id:   4,
			s:    "aaau",
			want: 1,
		},
	}

	for _, c := range cases {
		if res := maximumLength(c.s); res != c.want {
			fmt.Printf("Test %d failed, expected %d, got %d\n", c.id, c.want, res)
		}
	}
}

func maximumLength(s string) int {
	n := len(s)
	l, r := 1, n

	if !helper(s, n, l) {
		return -1
	}

	for l+1 < r {
		mid := (l + r) / 2
		if helper(s, n, mid) {
			l = mid
		} else {
			r = mid
		}
	}

	return l
}

func helper(s string, n, x int) bool {
	cnt := make([]int, 26)
	p := 0

	for i := 0; i < n; i++ {
		for s[p] != s[i] {
			p++
		}
		if i-p+1 >= x {
			cnt[s[i]-'a']++
		}
		if cnt[s[i]-'a'] > 2 {
			return true
		}
	}

	return false
}
