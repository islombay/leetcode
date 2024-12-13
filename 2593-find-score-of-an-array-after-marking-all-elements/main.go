package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	cases := []struct {
		id   int
		nums []int
		want int64
	}{
		{
			id:   1,
			nums: []int{2, 1, 3, 4, 5, 2},
			want: 7,
		},
		{
			id:   2,
			nums: []int{2, 3, 5, 1, 3, 2},
			want: 5,
		},
	}

	for _, c := range cases {
		if res := findScore2(c.nums); res != c.want {
			fmt.Printf("Test %d failed, expected %d, got %d\n", c.id, c.want, res)
		}
	}
}

type PriorityQueue [][2]int

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i][0] == pq[j][0] {
		return pq[i][1] < pq[j][1]
	}
	return pq[i][0] < pq[j][0]
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

func findScore(nums []int) int64 {
	idxMark := map[int]bool{}

	pq := &PriorityQueue{}
	heap.Init(pq)

	for i, n := range nums {
		heap.Push(pq, [2]int{n, i})
	}

	//for pq.Len() > 0 {
	//	item := heap.Pop(pq).([2]int)
	//	fmt.Println(item)
	//}
	//return 0

	var res int64 = 0

	for pq.Len() > 0 && len(idxMark) < len(nums) {
		top := heap.Pop(pq).([2]int)
		if _, ok := idxMark[top[1]]; ok {
			continue
		}

		idxMark[top[1]] = true
		if top[1] > 0 {
			idxMark[top[1]-1] = true
		}
		if top[1] < len(nums)-1 {
			idxMark[top[1]+1] = true
		}

		res += int64(top[0])
	}
	return res
}

func findScore2(nums []int) int64 {
	n := len(nums)
	sorted := make([][2]int, n)
	visited := make([]bool, n)
	var res int64 = 0

	for i, v := range nums {
		sorted[i] = [2]int{v, i}
	}

	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i][0] == sorted[j][0] {
			return sorted[i][1] < sorted[j][1]
		}
		return sorted[i][0] < sorted[j][0]
	})

	for _, top := range sorted {
		if visited[top[1]] == true {
			continue
		}

		visited[top[1]] = true
		if top[1] > 0 {
			visited[top[1]-1] = true
		}
		if top[1] < n-1 {
			visited[top[1]+1] = true
		}

		res += int64(top[0])
	}
	return res
}
