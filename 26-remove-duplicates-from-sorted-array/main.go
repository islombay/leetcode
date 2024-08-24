package main

import "fmt"

func main() {
	cases := []struct {
		name        string
		nums        []int
		result      int
		resultArray []int
	}{
		{
			name:        "1",
			nums:        []int{1, 2, 2},
			result:      2,
			resultArray: []int{1, 2},
		},
		{
			name:        "2",
			nums:        []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			resultArray: []int{0, 1, 2, 3, 4},
			result:      5,
		},
	}

	for _, c := range cases {
		if result := removeDuplicates(c.nums); result != c.result {
			fmt.Printf("- Test %s failed, expected result %d, expected array %v, got result %d, got array %v\n", c.name, c.result, c.resultArray, result, c.nums)
		} else if result == c.result {
			for i := range c.resultArray {
				if c.resultArray[i] != c.nums[i] {
					fmt.Printf("- Test %s failed, expected result %d, expected array %v, got result %d, got array %v\n", c.name, c.result, c.resultArray, result, c.nums)
				}
			}
			fmt.Printf("Test %s passed, expected result %d, expected array %v, got result %d, got array %v\n", c.name, c.result, c.resultArray, result, c.nums)
		}
	}
}

func removeDuplicates(nums []int) int {
	unique := make(map[int]int)
	removed := 0
	n := len(nums)

	for i := 0; i < len(nums); i++ {
		unique[nums[i]]++
		val, ok := unique[nums[i]]
		if ok && val > 1 {
			nums = append(nums[:i], nums[i+1:]...)
			removed++
			i--
		}
	}

	return n - removed
}
