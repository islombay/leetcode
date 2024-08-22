package main

import (
	"fmt"
	"reflect"
)

func main() {
	cases := []struct {
		nums   []int
		output int
		res    []int
	}{
		{
			nums:   []int{1, 1, 1, 2, 2, 3},
			res:    []int{1, 1, 2, 2, 3},
			output: 5,
		},
		{
			nums:   []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			res:    []int{0, 0, 1, 1, 2, 3, 3},
			output: 7,
		},
	}

	for _, c := range cases {
		res := removeDuplicates(c.nums)
		if res == c.output && reflect.DeepEqual(res, c.res) {
			fmt.Println("Test passed")
		} else {
			fmt.Println("Test failed", res, c.nums)
		}
	}
}

func removeDuplicates(nums []int) int {
	var (
		unique = make(map[int]int)
		count  = 0
		n      = len(nums)
		tmpLen = len(nums)
	)
	for i := 0; i < tmpLen; i++ {
		// count = length - total remove
		unique[nums[i]]++
		val, ok := unique[nums[i]]
		if ok && val > 2 {
			count++
			for j := i; j < tmpLen-1; j++ {
				nums[j] = nums[j+1]
			}
			tmpLen--
			i--
		}
	}
	return n - count
}
