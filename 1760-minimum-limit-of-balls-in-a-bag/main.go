package main

import "fmt"

func main() {
	cases := []struct {
		id     int
		nums   []int
		maxOps int
		want   int
	}{
		{
			id:     1,
			nums:   []int{9},
			maxOps: 2,
			want:   3,
		},
		{
			id:     2,
			nums:   []int{2, 4, 8, 2},
			maxOps: 4,
			want:   2,
		},
	}

	for _, c := range cases {
		if res := minimumSize(c.nums, c.maxOps); res != c.want {
			fmt.Printf("Test %d failed, expected %d, got %d\n", c.id, c.want, res)
		}
	}
}

func minimumSize(nums []int, maxOperations int) int {
	low, high := 1, 0
	for _, num := range nums {
		if num > high {
			high = num
		}
	}

	for low < high {
		mid := (low + high) / 2
		operationCount := 0
		for _, num := range nums {
			operationCount += (num - 1) / mid
		}
		if operationCount <= maxOperations {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return high
}
