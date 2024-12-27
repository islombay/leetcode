package main

import "fmt"

func main() {
	cases := []struct {
		id   int
		v    []int
		want int
	}{
		{
			id:   1,
			v:    []int{8, 1, 5, 2, 6},
			want: 11,
		},
		{
			id:   2,
			v:    []int{1, 2},
			want: 2,
		},
		{
			id:   3,
			v:    []int{1, 3, 5},
			want: 7,
		},
		{
			id:   4,
			v:    []int{1, 2, 2},
			want: 3,
		},
	}

	for _, c := range cases {
		if res := maxScoreSightseeingPair(c.v); res != c.want {
			fmt.Printf("Test %d failed, expected %d, got %d\n", c.id, c.want, res)
		}
	}
}

func maxScoreSightseeingPair(values []int) int {
	maxFirstIdx := 0
	for i, n := range values {
		if (n-i > values[maxFirstIdx]-maxFirstIdx) || (n-i == values[maxFirstIdx]-maxFirstIdx && n > values[maxFirstIdx]) {
			maxFirstIdx = i
		}
	}

	maxSecondIdx := 0
	if maxSecondIdx == maxFirstIdx {
		maxSecondIdx++
	}
	for i, n := range values {
		if i != maxFirstIdx && n-i > values[maxSecondIdx]-maxSecondIdx {
			maxSecondIdx = i
		}
	}

	if maxSecondIdx < maxFirstIdx {
		maxSecondIdx, maxFirstIdx = maxFirstIdx, maxSecondIdx
	}

	return (values[maxSecondIdx] - maxSecondIdx) + (values[maxFirstIdx] + maxFirstIdx)
}

// COpied but working solution
func maxScoreSightseeingPair2(values []int) int {
	bestLeft := values[0] + 0
	maxScore := 0

	for i := 1; i < len(values); i++ {
		maxScore = max(maxScore, bestLeft+values[i]-i)
		bestLeft = max(bestLeft, values[i]+i)
	}
	return maxScore
}
