package main

import "fmt"

func main() {
	cases := []struct {
		name   string
		board  [][]byte
		result bool
	}{
		{
			name: "#1",
			board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '.', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			result: true,
		},
		{
			name: "#2",
			board: [][]byte{
				{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			result: false,
		},
		{
			name: "#3",
			board: [][]byte{
				{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
				{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
				{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
				{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
				{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
				{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
				{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
			},
			result: false,
		},
	}

	for _, c := range cases {
		result := isValidSudoku(c.board)
		if result != c.result {
			fmt.Printf("Test %s failed, expected %b, got %b\n", c.name, c.result, result)
		} else {
			fmt.Printf("Test finished %s\n", c.name)
		}
	}
}

func isValidSudoku(board [][]byte) bool {
	// check each column and row

	var (
		mapList    = make(map[byte]int)
		rowLen     = len(board[0])
		posY, posX = 3, 3
	)

	for _, row := range board {
		for _, cell := range row {
			if cell != '.' {
				if val, ok := mapList[cell]; ok && val >= 1 {
					return false
				}
				mapList[cell]++
			}
		}
		mapList = make(map[byte]int)
	}

	for i := 0; i < rowLen; i++ {
		for _, row := range board {
			if row[i] != '.' {
				if val, ok := mapList[row[i]]; ok && val >= 1 {
					return false
				}
				mapList[row[i]]++
			}
		}
		mapList = make(map[byte]int)
	}

	for {
		for _, row := range board[posY-3 : posY] {
			for _, cell := range row[posX-3 : posX] {
				if cell != '.' {
					if val, ok := mapList[cell]; ok && val >= 1 {
						return false
					}
					mapList[cell]++
				}
			}
		}
		posX += 3
		mapList = make(map[byte]int)
		if posX == 12 && posY != 9 {
			posX = 3
			posY += 3
		} else if posX == 12 && posY == 9 {
			break
		}
	}
	return true
}
