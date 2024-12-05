package main

import "fmt"

func main() {
	cases := []struct {
		id     int
		n1, n2 []int
		want   []int
	}{
		{
			id:   1,
			n1:   []int{4, 1, 2},
			n2:   []int{1, 3, 4, 2},
			want: []int{-1, 3, -1},
		},
		{
			id:   2,
			n1:   []int{2, 4},
			n2:   []int{1, 2, 3, 4},
			want: []int{3, -1},
		},
	}

	for _, c := range cases {
		res := nextGreaterElement(c.n1, c.n2)
		if len(res) != len(c.want) {
			fmt.Printf("Test %d failed len\n", c.id)
			continue
		}
		for i := 0; i < len(res); i++ {
			if res[i] != c.want[i] {
				fmt.Printf("Test %d failed\n", c.id)
				break
			}
		}
	}
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := []int{}
	m := make(map[int]int)

	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums2[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			m[nums2[i]] = -1
		} else {
			m[nums2[i]] = stack[len(stack)-1]
		}

		stack = append(stack, nums2[i])
	}

	res := make([]int, len(nums1))
	for i, num := range nums1 {
		res[i] = m[num]
	}
	return res
}
