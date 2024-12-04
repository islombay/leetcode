package main

import "fmt"

func main() {
	cases := []struct {
		id   int
		grid [][]int
		want int
	}{
		{
			id:   1,
			grid: [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}},
			want: 16,
		},
		{
			id:   2,
			grid: [][]int{{1}},
			want: 4,
		},
		{
			id:   3,
			grid: [][]int{{1, 0}},
			want: 4,
		},
	}

	for _, c := range cases {
		if res := islandPerimeter(c.grid); res != c.want {
			fmt.Printf("Test %d failed, expected %d, got %d\n", c.id, c.want, res)
		}
	}
}

func islandPerimeter(grid [][]int) int {
	perimeter := 0
	for i := 0; i < len(grid); i++ {
		prevJ := -1
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				perimeter += 4

				if prevJ != -1 && j-1 == prevJ {
					perimeter -= 2
				}

				if i != 0 && grid[i-1][j] == 1 {
					perimeter -= 2
				}

				prevJ = j
			}
		}
	}

	return perimeter
}
