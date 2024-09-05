package main

import "fmt"

func main() {
	cases := []struct {
		name   string
		nums   []int
		k      int
		result bool
	}{
		{
			name:   "1",
			nums:   []int{1, 2, 3, 1},
			k:      3,
			result: true,
		},
		{
			name:   "2",
			nums:   []int{1, 0, 1, 1},
			k:      1,
			result: true,
		},
		{
			name:   "3",
			nums:   []int{1, 2, 3, 1, 2, 3},
			k:      2,
			result: false,
		},
	}

	for _, c := range cases {
		if res := containsNearbyDuplicate(c.nums, c.k); res != c.result {
			fmt.Printf("- Test %s failed, expected %t, got %t\n", c.name, c.result, res)
		} else {
			fmt.Printf("Test %s passed, expected %t, got %t\n", c.name, c.result, res)
		}
	}
}

func containsNearbyDuplicate(nums []int, k int) bool {
	unq := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := unq[nums[i]]; ok {
			return true
		}

		unq[nums[i]]++

		if len(unq) > k {
			delete(unq, nums[i-k])
		}
	}
	return false
}
