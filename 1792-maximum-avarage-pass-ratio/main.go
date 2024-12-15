package main

import (
	"container/heap"
	"fmt"
)

func main() {
	cases := []struct {
		id            int
		classes       [][]int
		extraStudents int
		want          float64
	}{
		{
			id:            1,
			classes:       [][]int{{1, 2}, {3, 5}, {2, 2}},
			extraStudents: 2,
			want:          0.78333,
		},
		{
			id:            2,
			classes:       [][]int{{2, 4}, {3, 9}, {4, 5}, {2, 10}},
			extraStudents: 4,
			want:          0.53485,
		},
	}

	for _, c := range cases {
		if res := maxAverageRatio(c.classes, c.extraStudents); res != c.want {
			fmt.Printf("Test %d failed, expected %f, got %f\n", c.id, c.want, res)
		}
	}
}

type PriorityQueue [][2]int

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return float64(pq[i][0]+1)/float64(pq[i][1]+1)-float64(pq[i][0])/float64(pq[i][1]) > float64(pq[j][0]+1)/float64(pq[j][1]+1)-float64(pq[j][0])/float64(pq[j][1])
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.([2]int))
}

func (pq *PriorityQueue) Pop() any {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	pq := &PriorityQueue{}
	heap.Init(pq)

	for _, cl := range classes {
		heap.Push(pq, [2]int{cl[0], cl[1]})
	}

	for extraStudents > 0 {
		top := heap.Pop(pq).([2]int)

		top[0]++
		top[1]++
		heap.Push(pq, top)
		extraStudents--
	}

	var res float64
	for _, cl := range *pq {
		res += float64(cl[0]) / float64(cl[1])
	}

	return res / float64(len(classes))
}
