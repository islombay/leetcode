package main

import "fmt"

type NumArray struct {
	nums []int
	pref []int
}

func Constructor(nums []int) NumArray {
	pref := make([]int, len(nums)+1)
	for i, n := range nums {
		pref[i+1] = pref[i] + n
	}
	return NumArray{nums: nums, pref: pref}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.pref[right+1] - this.pref[left]
}

func main() {
	cases := []struct {
		id     int
		nums   []int
		values []struct {
			left   int
			right  int
			answer int
		}
	}{
		{
			id:   1,
			nums: []int{-2, 0, 3, -5, 2, -1},
			values: []struct {
				left   int
				right  int
				answer int
			}{
				{
					left:   0,
					right:  2,
					answer: 1,
				},
				{
					left:   2,
					right:  5,
					answer: -1,
				},
				{
					left:   0,
					right:  5,
					answer: -3,
				},
			},
		},
	}

	for _, c := range cases {
		obj := Constructor(c.nums)
		for i, v := range c.values {
			if ans := obj.SumRange(v.left, v.right); ans != v.answer {
				fmt.Printf("Test %d failed, value id %d, expected %d, got %d\n", c.id, i, v.answer, ans)
				break
			}
		}
	}
	fmt.Println("Finish")
}
