package main

import "fmt"

func main() {
	cases := []struct {
		id            int
		start, target string
		want          bool
	}{
		{
			id:     1,
			start:  "_L__R__R_",
			target: "L______RR",
			want:   true,
		},
		{
			id:     2,
			start:  "R_L_",
			target: "__LR",
			want:   false,
		},
		{
			id:     3,
			start:  "_R",
			target: "R_",
			want:   false,
		},
	}

	for _, c := range cases {
		if res := canChange(c.start, c.target); res != c.want {
			fmt.Printf("Test %d failed, expected %t, got %t\n", c.id, c.want, res)
		}
	}
}

func canChange(start string, target string) bool {
	if start == target {
		return true
	}
	waitL, waitR := 0, 0

	for i := 0; i < len(start); i++ {
		curr := start[i]
		goal := target[i]

		if curr == 'R' {
			if waitL > 0 {
				return false
			}
			waitR++
		}
		if goal == 'L' {
			if waitR > 0 {
				return false
			}
			waitL++
		}
		if goal == 'R' {
			if waitR == 0 {
				return false
			}
			waitR--
		}
		if curr == 'L' {
			if waitL == 0 {
				return false
			}
			waitL--
		}
	}
	return waitL == 0 && waitR == 0
}
