package main

import "fmt"

func main() {
	cases := []struct {
		id      int
		numRows int
		want    [][]int
	}{
		{
			id:      1,
			numRows: 1,
			want:    [][]int{{1}},
		},
		{
			id:      2,
			numRows: 5,
			want:    [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
	}

	for _, c := range cases {
		res := generate(c.numRows)
		if len(res) != len(c.want) {
			fmt.Printf("Test %d failed\n", c.id)
			continue
		}
		for i, arr := range c.want {
			if len(res[i]) != len(arr) {
				fmt.Printf("Test %d failed\n", c.id)
				break
			}
			for j := range arr {
				if arr[j] != res[i][j] {
					fmt.Printf("Test %d failed\n", c.id)
					break
				}
			}
		}
	}
	fmt.Println("Finished")
}
func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	res[0] = []int{1}

	for i := 1; i < numRows; i++ {
		res[i] = make([]int, len(res[i-1])+1)
		res[i][0] = 1
		if len(res[i-1]) == 1 {
			res[i][1] = 1
			continue
		}
		for j := 1; j < len(res[i-1]); j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
		res[i][len(res[i])-1] = 1
	}
	return res
}
