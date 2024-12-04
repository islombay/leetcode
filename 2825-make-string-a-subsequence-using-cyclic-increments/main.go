package main

import "fmt"

func main() {
	cases := []struct {
		id             int
		source, target string
		want           bool
	}{
		{
			id:     1,
			source: "abc",
			target: "ad",
			want:   true,
		},
		{
			id:     2,
			source: "zc",
			target: "ad",
			want:   true,
		},
		{
			id:     3,
			source: "ab",
			target: "d",
			want:   false,
		},
	}

	for _, c := range cases {
		if res := canMakeSubsequence(c.source, c.target); res != c.want {
			fmt.Printf("Test %d failed, expected %t, got %t\n", c.id, c.want, res)
		}
	}
}

func canMakeSubsequence(source string, target string) bool {
	targetIdx, targetLen := 0, len(target)

	for _, currChar := range source {
		if targetIdx < targetLen && (int(target[targetIdx])-int(currChar)+26)%26 <= 1 {
			targetIdx++
		}
	}
	return targetIdx == targetLen
}
