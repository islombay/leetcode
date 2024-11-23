package main

import "fmt"

func main() {
	cases := []struct {
		name        string
		nums        []int
		val         int
		resultArray []int
		result      int
	}{
		{
			name:        "1",
			nums:        []int{3, 2, 2, 3},
			val:         3,
			resultArray: []int{2, 2},
			result:      2,
		},
		{
			name:        "2",
			nums:        []int{0, 1, 2, 2, 3, 0, 4, 2},
			resultArray: []int{0, 1, 3, 0, 4},
			val:         2,
			result:      5,
		},
	}

	for _, c := range cases {
		res := removeElement(c.nums, c.val)
		if res != c.result {
			fmt.Printf("Test %s failed, expected %d, got %d, expected array %v, got array %v\n", c.name, c.result, res, c.resultArray, c.nums)
		} else {
			for i := range c.resultArray {
				if c.resultArray[i] != c.nums[i] {
					fmt.Printf("Test %s failed, expected %d, got %d, expected array %v, got array %v\n", c.name, c.result, res, c.resultArray, c.nums)
				}
			}
		}
	}
	fmt.Println("Finished")
}

func removeElement(nums []int, val int) int {
	arrIndex := 0
	for _, n := range nums {
		if n != val {
			nums[arrIndex] = n
			arrIndex++
		}
	}
	return arrIndex
}
