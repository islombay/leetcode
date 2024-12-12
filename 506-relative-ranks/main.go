package main

import (
	"container/heap"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(findRelativeRanks([]int{5, 4, 3, 2, 1}))
}

type PriorityQueue [][2]int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// Max-heap: larger scores have higher priority
	return pq[i][0] > pq[j][0]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.([2]int))
}

func (pq *PriorityQueue) Pop() any {
	n := len(*pq)
	res := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return res
}

func findRelativeRanks(score []int) []string {
	pq := &PriorityQueue{}
	heap.Init(pq)

	for i, v := range score {
		heap.Push(pq, [2]int{v, i})
	}

	res := make([]string, len(score))
	rank := 1
	for pq.Len() > 0 {
		top := heap.Pop(pq).([2]int)
		if rank == 1 {
			res[top[1]] = "Gold Medal"
		} else if rank == 2 {
			res[top[1]] = "Silver Medal"
		} else if rank == 3 {
			res[top[1]] = "Bronze Medal"
		} else {
			res[top[1]] = strconv.Itoa(rank)
		}
		rank++
	}
	return res
}
