package main

import "fmt"

func main() {
	cases := []struct {
		id     int
		nums   []int
		target int
		want   bool
	}{
		{
			id:     1,
			nums:   []int{2, 5, 6, 0, 0, 1, 2},
			target: 0,
			want:   true,
		},
		{
			id:     2,
			nums:   []int{2, 5, 6, 0, 0, 1, 2},
			target: 3,
			want:   false,
		},
		{
			id:     3,
			nums:   []int{1, 0, 1, 1, 1},
			target: 0,
			want:   true,
		},
		{
			id:     4,
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1},
			target: 2,
			want:   true,
		},
	}

	for _, c := range cases {
		if res := search(c.nums, c.target); res != c.want {
			fmt.Printf("Test %d failed, expected %t, got %t\n", c.id, c.want, res)
		}
	}
}

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return true
		} else if target < nums[mid] {
			if nums[mid] >= nums[l] && mid != l {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else if target > nums[mid] {
			if nums[mid] <= nums[r] && mid != r {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return false
}
