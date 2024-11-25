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
			nums: []int{2, 2, 1},
			want: 1,
		},
		{
			id:   2,
			nums: []int{4, 1, 2, 1, 2},
			want: 4,
		},
	}

	for _, c := range cases {
		res := singleNumber(c.nums)
		if res != c.want {
			fmt.Printf("Test %d failed, expected %d, got %d\n", c.id, c.want, res)
		}
	}
	fmt.Println("Finished")
}

func singleNumber(nums []int) int {
	result := 0
	for _, v := range nums {
		result ^= v
	}
	return result
}
