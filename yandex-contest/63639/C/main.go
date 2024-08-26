package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var N int
	fmt.Fscan(reader, &N)

	array := make([][]int, N)
	intersecretArrays := make(map[int]int)

	for i := 0; i < N; i++ {
		var X, Y int
		fmt.Fscan(reader, &X, &Y)

		array[i] = []int{X, Y}
	}

Out:
	for i := 0; i < N; i++ {
		j := i + 1

		for ; j < N; j++ {
			if array[i][0] != array[j][0] && array[i][1] == array[j][1] {
				intersecretArrays[array[i][0]] = array[i][1]
				intersecretArrays[array[j][0]] = array[j][1]
				continue Out
			} else if array[i][0] <= array[j][0] && array[i][1] > array[j][1] ||
				array[i][0] >= array[j][0] && array[i][1] < array[j][1] {
				intersecretArrays[array[i][0]] = array[i][1]
				intersecretArrays[array[j][0]] = array[j][1]
			}
		}

	}

	fmt.Println(N - len(intersecretArrays))
}
