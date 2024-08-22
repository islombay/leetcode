package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n) // Read the first input as a number

	// Create a slice to store the array of integers
	arr := make([]int, n)
	list := make(map[int]int)
	count := 0

	// Read the array of integers
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		list[arr[i]]++
		val, ok := list[arr[i]]
		if ok && val >= 2 {
			count--
		} else {
			count++
		}
	}

	fmt.Println(count)
}
