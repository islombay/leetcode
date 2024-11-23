package main

import "fmt"

func main() {
	cases := []struct {
		id     int
		nums   []int
		target int
		result int
	}{
		{
			id:     1,
			nums:   []int{1, 3, 5, 6},
			target: 5,
			result: 2,
		},
		{
			id:     2,
			nums:   []int{1, 3, 5, 6},
			target: 2,
			result: 1,
		},
		{
			id:     3,
			nums:   []int{1, 3, 5, 6},
			target: 7,
			result: 4,
		},
	}

	for _, c := range cases {
		if res := searchInsert(c.nums, c.target); res != c.result {
			fmt.Printf("Test %d failed, expected %d, got %d\n", c.id, c.result, res)
		}
	}
	fmt.Println("Finished")
}

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
