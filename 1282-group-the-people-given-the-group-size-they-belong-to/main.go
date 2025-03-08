package main

import "fmt"

func main() {
	groupSize := []int{3, 3, 3, 3, 3, 1, 3}
	fmt.Println(groupThePeople(groupSize))
}

func groupThePeople(groupSizes []int) [][]int {
	groupMap := map[int][]int{}
	result := make([][]int, 0, len(groupSizes))
	for i, e := range groupSizes {
		_, exists := groupMap[e]
		if !exists {
			groupMap[e] = make([]int, 0, e)
		}

		groupMap[e] = append(groupMap[e], i)

		if len(groupMap[e]) == e {
			result = append(result, groupMap[e])
			delete(groupMap, e)
		}
	}

	return result
}
