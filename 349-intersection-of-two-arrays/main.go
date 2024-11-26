package main

import (
	"fmt"
	"sort"
)

func main() {
	cases := []struct {
		id  int
		n1  []int
		n2  []int
		res []int
	}{
		{
			id:  1,
			n1:  []int{1, 2, 2, 1},
			n2:  []int{2, 2},
			res: []int{2},
		},
		{
			id:  2,
			n1:  []int{4, 9, 5},
			n2:  []int{9, 4, 9, 8, 4},
			res: []int{9, 4},
		},
	}

	for _, c := range cases {
		res := intersection(c.n1, c.n2)
		if len(res) != len(c.res) {
			fmt.Printf("Test %d failed, len\n", c.id)
			continue
		}

		sort.Ints(res)
		sort.Ints(c.res)
		for i := 0; i < len(res); i++ {
			if res[i] != c.res[i] {
				fmt.Printf("Test %d failed\n", c.id)
				break
			}
		}
	}
	fmt.Println("Finished")
}

func intersection(nums1 []int, nums2 []int) []int {
	n1 := make(map[int]bool)
	resMap := make(map[int]int)

	for _, v := range nums1 {
		n1[v] = true
	}
	for _, v := range nums2 {
		if n1[v] {
			resMap[v]++
		}
	}

	resArray := []int{}
	for k, _ := range resMap {
		resArray = append(resArray, k)
	}
	return resArray
}
