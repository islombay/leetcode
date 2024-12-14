package main

func shiftGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])
	res := make([][]int, m)

	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			x := (k / n) % m
			y := k % n

			res[x][y] = grid[i][j]
			k++
		}
	}
	return res
}
