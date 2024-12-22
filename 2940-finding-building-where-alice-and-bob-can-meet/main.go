package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	cases := []struct {
		id      int
		heights []int
		queries [][]int
		want    []int
	}{
		{
			id:      1,
			heights: []int{6, 4, 8, 5, 2, 7},
			queries: [][]int{{0, 1}, {0, 3}, {2, 4}, {3, 4}, {2, 2}},
			want:    []int{2, 5, -1, 5, 2},
		},
		{
			id:      2,
			heights: []int{3, 1, 2, 4},
			queries: [][]int{{0, 1}, {1, 0}},
			want:    []int{3, 3},
		},
		{
			id:      3,
			heights: []int{3, 4, 1, 2},
			queries: [][]int{{0, 2}},
			want:    []int{-1},
		},
	}

	for _, c := range cases {
		res := leftmostBuildingQueries(c.heights, c.queries)
		if !reflect.DeepEqual(res, c.want) {
			fmt.Printf("Test %d failed\n", c.id)
		}
	}
}

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	ll := len(queries)
	indies := make([]int, ll)
	for i := range ll {
		if queries[i][0] > queries[i][1] {
			queries[i][0], queries[i][1] = queries[i][1], queries[i][0]
		}
		indies[i] = i
	}
	sort.Slice(indies, func(i, j int) bool {
		return queries[indies[i]][1] > queries[indies[j]][1]
	})
	ans := make([]int, ll)
	for i := range ll {
		ans[i] = -1
	}

	l := len(heights)
	stack := make([]int, l)
	stackIndex := l - 1
	stack[stackIndex] = l - 1
	cur := l - 2

	for _, i := range indies {
		a, b := queries[i][0], queries[i][1]
		if a == b || heights[a] < heights[b] {
			ans[i] = b
			continue
		}
		if b == l-1 {
			continue
		}
		target := max(heights[a], heights[b])
		for ; cur > b; cur-- {
			for ; stackIndex != l && heights[cur] >= heights[stack[stackIndex]]; stackIndex++ {
			}
			stackIndex--
			stack[stackIndex] = cur
		}
		if idx := sort.Search(l-stackIndex, func(i int) bool {
			return heights[stack[i+stackIndex]] > target
		}); idx != l-stackIndex {
			ans[i] = stack[idx+stackIndex]
		}
	}

	return ans
}
