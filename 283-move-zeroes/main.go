package main

import "fmt"

func main() {
	cases := []struct {
		id   int
		nums []int
		want []int
	}{
		{
			id:   1,
			nums: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			id:   2,
			nums: []int{0},
			want: []int{0},
		},
		{
			id:   3,
			nums: []int{0, 1},
			want: []int{1, 0},
		},
	}

	for _, c := range cases {
		moveZeroes(c.nums)
		for i := range c.want {
			if c.want[i] != c.nums[i] {
				fmt.Printf("Test %d failed, index %d\n", c.id, i)
			}
		}
	}
	fmt.Println("Finished")
}

func moveZeroes(nums []int) {
	idx := 0
	found := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			found++
			continue
		}
		nums[idx] = nums[i]
		idx++
	}

	for i := len(nums) - found; i < len(nums); i++ {
		nums[i] = 0
	}
}
