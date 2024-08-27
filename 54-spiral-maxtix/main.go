package main

import "fmt"

func main() {
	cases := []struct {
		name   string
		matrix [][]int
		result []int
	}{
		{
			name: "1",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			result: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "2",
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			result: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			name: "3",
			matrix: [][]int{
				{2, 3, 4},
				{5, 6, 7},
				{8, 9, 10},
				{11, 12, 13},
				{14, 15, 16},
			},
			result: []int{2, 3, 4, 7, 10, 13, 16, 15, 14, 11, 8, 5, 6, 9, 12},
		},
	}

out:
	for _, c := range cases {
		r := spiralOrder(c.matrix)
		if len(r) == len(c.result) {
			for i := range r {
				if r[i] != c.result[i] {
					fmt.Printf("- Test %s failed, expected %v, got %v\n", c.name, c.result, r)
					continue out

				}
			}
			fmt.Printf("Test %s passed, expected %v, got %v\n", c.name, c.result, r)
		} else {
			fmt.Printf("- Test %s failed, expected %v, got %v\n", c.name, c.result, r)
		}
	}
}

func spiralOrder(matrix [][]int) []int {
	var (
		startM = 0
		endM   = len(matrix)

		startN = 0
		endN   = len(matrix[0])

		result []int
	)

	if endM == 1 {
		return matrix[0]
	}
	if endN == 1 {
		for i := range matrix {
			result = append(result, matrix[i][0])
		}
		return result
	}

	for startN < endN && startM < endM {
		// get upper numbers
		for i := startN; i < endN; i++ {
			result = append(result, matrix[startM][i])
		}

		// get right numbers
		for i := startM + 1; i < endM; i++ {
			result = append(result, matrix[i][endN-1])
		}

		// get bottom numbers
		for i := endN - 2; i >= startN; i-- {
			if endM-1 != startM {
				result = append(result, matrix[endM-1][i])
			}
		}

		// get left numbers
		for i := endM - 2; i >= startM+1; i-- {
			if endN-1 != startN {
				result = append(result, matrix[i][startN])
			}
		}

		// change coordinates
		startM++
		endM--
		startN++
		endN--
	}

	return result
}
