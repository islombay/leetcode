package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		id   int
		nums []int
		want int64
	}{
		{
			id:   1,
			nums: []int{5, 4, 2, 4},
			want: 8,
		},
		{
			id: 2,
			nums: []int{1,2,3},
			want: 6,
,		},
	}

	for _, c := range cases {
		if res := continuousSubarrays(c.nums); res != c.want {
			fmt.Printf("Test %d failed, expected %d, got %d\n", c.id, c.want, res)
		}
	}
}

func continuousSubarrays(nums []int) int64 {

	var ans int64

	maxNum := nums[0]
	minNum := nums[0]

	for l, r := 0, 0; r < len(nums); r++ {
		if abs(nums[r] - maxNum) <= 2 && abs(nums[r] - minNum) <= 2{
			ans += int64(r - l + 1)
			maxNum = max(maxNum, nums[r])
			minNum = min(minNum, nums[r])
			continue
		}
		maxNum = nums[r]
		minNum = nums[r]

		for j := r; j >= l; j-- {
			if abs(nums[r] - nums[j]) > 2 {
				l = j+1
				break
			}
			maxNum = max(maxNum, nums[j])
			minNum = min(minNum, nums[j])
		}
		ans += int64(r - l + 1)
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}