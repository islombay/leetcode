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
			nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
			want: []int{5, 6},
		},
		{
			id:   2,
			nums: []int{1, 1},
			want: []int{2},
		},
	}

	for _, c := range cases {
		res := findDisappearedNumbers(c.nums)
		if len(res) != len(c.want) {
			fmt.Printf("Test %d failed, len\n", c.id)
			continue
		}

		for i := 0; i < len(res); i++ {
			if res[i] != c.want[i] {
				fmt.Printf("Test %d failed\n", c.id)
				break
			}
		}
	}
	fmt.Println("Finished")
}

func findDisappearedNumbers(nums []int) []int {
	res := make([]int, len(nums))
	resArray := []int{}
	for _, n := range nums {
		res[n-1] = n
	}

	for i, n := range res {
		if n == 0 {
			resArray = append(resArray, i+1)
		}
	}
	return resArray
}
