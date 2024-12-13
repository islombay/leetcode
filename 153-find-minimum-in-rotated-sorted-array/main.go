package main

import "fmt"

func main() {
	cases := []struct {
		id   int
		nums []int
		want int
	}{
		{
			id:   1,
			nums: []int{3, 4, 5, 1, 2},
			want: 1,
		},
		{
			id:   2,
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			want: 0,
		},
		{
			id:   3,
			nums: []int{11, 13, 15, 17},
			want: 11,
		},
	}

	for _, c := range cases {
		if res := findMin(c.nums); res != c.want {
			fmt.Printf("Test %d failed, expected %d, got %d\n", c.id, c.want, res)
		}
	}
}

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	min := nums[0]

	for l <= r {
		mid := l + (r-l)/2
		if min > nums[mid] {
			min = mid
		}
		if mid > 0 && nums[mid] < nums[mid-1] {
			return nums[mid]
		}
		if mid < len(nums)-1 && nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return min
}
