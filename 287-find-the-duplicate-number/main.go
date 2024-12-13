package main

import "fmt"

func main() {
	cases := []struct {
		id   int
		nums []int
		n    int
	}{
		{
			id:   1,
			nums: []int{1, 2, 3, 2},
			n:    2,
		},
		{
			id:   2,
			nums: []int{3, 1, 3, 4, 2},
			n:    3,
		},
	}

	for _, c := range cases {
		if res := findDuplicate(c.nums); res != c.n {
			fmt.Printf("Test %d failed, expected %d, got %d\n", c.id, c.n, res)
		}
	}
}

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			break
		}
	}

	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
