package main

import "fmt"

func main() {
	cases := []struct {
		id       int
		rowIndex int
		want     []int
	}{
		{
			id:       1,
			rowIndex: 1,
			want:     []int{1, 1},
		},
		{
			id:       2,
			rowIndex: 0,
			want:     []int{1},
		},
		{
			id:       3,
			rowIndex: 6,
			want:     []int{1, 6, 15, 20, 15, 6, 1},
		},
	}

	for _, c := range cases {
		res := getRow(c.rowIndex)
		if len(res) != len(c.want) {
			fmt.Printf("Test %d failed, len\n", c.id)
			continue
		}
		for i := 0; i < len(res); i++ {
			if res[i] != c.want[i] {
				fmt.Printf("Test %d failed, index %d\n", c.id, c.rowIndex)
				break
			}
		}
	}
	fmt.Println("Finished")
}

func getRow(rowIndex int) []int {
	rows := make([]int, rowIndex+1)
	rows[0] = 1
	rows[rowIndex] = 1
	mid := (rowIndex + 1) / 2

	for i := 1; i <= mid; i++ {
		rows[i] = rows[i-1] * (rowIndex - i + 1) / i
		rows[rowIndex+1-(i+1)] = rows[i]
	}
	return rows
}
