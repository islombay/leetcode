package main

import "fmt"

func main() {
	cases := []struct {
		matrix [][]int
	}{
		//{
		//	matrix: [][]int{
		//		{1, 1, 1},
		//		{1, 0, 1},
		//		{1, 1, 1},
		//	},
		//},
		{
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 0, 7, 8},
				{0, 10, 11, 12},
				{13, 14, 15, 0},
			},
		},
	}

	for _, c := range cases {
		setZeroes(c.matrix)
	}
}

type Pos struct {
	X, Y int
}

func setZeroes(matrix [][]int) {
	var (
		m = len(matrix)
		n = len(matrix[0])

		changeList = make(map[Pos]int)
		// 0 -> first time -> setZero
		// 1 -> changed to zero -> do not do
		// 2 -> was zero -> setZero
	)

	if m == 1 && n == 1 {
		return
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				if val, ok := changeList[Pos{X: i, Y: j}]; (ok && val == 1) || (!ok && val == 0) {
					fmt.Printf("I: %d, J: %d, ok: %t, val %d\n", i, j, ok, val)
					setZero(matrix, i, j, m, n, changeList)
					fmt.Println(matrix)
				}
			}
		}
	}

	return
}

func setZero(matrix [][]int, m, n, mLen, nLen int, changeList map[Pos]int) {
	for i := 0; i < mLen; i++ {
		for j := 0; j < nLen; j++ {
			if i == m {
				if matrix[i][j] == 0 {
					if val, ok := changeList[Pos{X: i, Y: j}]; ok && val != 2 {
						changeList[Pos{X: i, Y: j}] = 1
					}
				} else {
					changeList[Pos{X: i, Y: j}] = 2
				}
				matrix[i][j] = 0
			}
			if j == n {
				if matrix[i][j] == 0 {
					if val, ok := changeList[Pos{X: i, Y: j}]; ok && val != 2 {
						changeList[Pos{X: i, Y: j}] = 1
					}
				} else {
					changeList[Pos{X: i, Y: j}] = 2
				}
				matrix[i][j] = 0
			}
		}
	}
}
