package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		id   int
		num  string
		want bool
	}{
		{
			id:   1,
			num:  "1234",
			want: false,
		},
		{
			id:   2,
			num:  "24123",
			want: true,
		},
		{
			id:   3,
			num:  "123",
			want: false,
		},
	}

	for _, c := range cases {
		if res := isBalanced(c.num); res != c.want {
			fmt.Printf("Test %d, failed, expected %t, got %t\n", c.id, c.want, res)
		}
	}
}

func isBalanced(num string) bool {
	total := 0
	sign := 1

	for i, _ := range num {
		total += int(num[i]-'0') * sign
		sign = -sign
	}
	return total == 0
}
