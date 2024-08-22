package main

import "fmt"

func main() {
	arr := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Scan(&arr[i])
		if i == 1 {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
		}
		if i == 2 {
			if arr[i] < arr[i-2] {
				arr[i], arr[i-2] = arr[i-2], arr[i]
			}
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
		}
	}
	fmt.Println(arr[1])
}
