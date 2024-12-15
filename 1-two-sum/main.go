package main

import (
	"fmt"
)

func main() {
	tests := []struct {
		id     int
		nums   []int
		target int
		res    []int
	}{
		{
			id:     1,
			nums:   []int{2, 7, 11, 15},
			target: 9,
			res:    []int{0, 1},
		},
		{
			id:     2,
			nums:   []int{3, 2, 4},
			target: 6,
			res:    []int{1, 2},
		},
		{
			id:     3,
			nums:   []int{3, 3},
			target: 6,
			res:    []int{0, 1},
		},
	}

	for _, test := range tests {
		res := twoSum(test.nums, test.target)
		if len(res) != len(test.res) {
			fmt.Printf("Test %d failed, len mismatch\n\n", test.id)
			continue
		}
		for i := 0; i < len(res); i++ {
			if res[i] != test.res[i] {
				fmt.Printf("Test %d failed, result mismatch\nExpected: %v\nGot: %v\n\n", test.id, test.res, res)
				break
			}
		}
	}
	fmt.Println("Finished")
}

func twoSum(nums []int, target int) []int {
	nKeys := map[int]int{}
	for i, num := range nums {
		if idx, ok := nKeys[target-num]; ok {
			return []int{idx, i}
		}
		nKeys[num] = i
	}
	return []int{0, 1}
}
