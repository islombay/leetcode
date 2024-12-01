package main

import (
	"fmt"
	"sort"
)

func main() {
	cases := []struct {
		id   int
		arr  []int
		want bool
	}{
		{
			id:   1,
			arr:  []int{10, 2, 5, 3},
			want: true,
		},
		{
			id:   2,
			arr:  []int{3, 1, 7, 11},
			want: false,
		},
		{
			id:   3,
			arr:  []int{-10, 12, -20, -8, 15},
			want: true,
		},
		{
			id:   4,
			arr:  []int{-2, 0, 10, -19, 4, 6, -8},
			want: false,
		},
		{
			id:   5,
			arr:  []int{0, 0},
			want: true,
		},
	}

	for _, c := range cases {
		if res := checkIfExist(c.arr); res != c.want {
			fmt.Printf("Test %d failed, expected %t, got %t\n", c.id, c.want, res)
		}
	}
}

func checkIfExist(arr []int) bool {
	sort.Ints(arr)
	found := map[int]int{}
	for i, n := range arr {
		found[n] = i
	}

	for i, n := range arr {
		// if _, ok := found[n]; ok {
		//     return true
		// }
		if val, ok := found[n*2]; ok {
			if n == 0 && val != i {
				return true
			} else if n != 0 {
				return true
			}
		}
	}
	return false
}
