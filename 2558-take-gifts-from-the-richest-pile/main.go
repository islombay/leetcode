package main

import (
	"container/heap"
	"math"
)

type Queue []int

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Less(i, j int) bool {
	return q[i] > q[j]
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *Queue) Push(x any) {
	*q = append(*q, x.(int))
}

func (q *Queue) Pop() any {
	n := len(*q)
	res := (*q)[n-1]
	*q = (*q)[:n-1]
	return res
}

func pickGifts(gifts []int, k int) int64 {
	h := Queue(gifts)
	heap.Init(&h)

	for _ = range k {
		current := heap.Pop(&h).(int)
		gift := int(math.Sqrt(float64(current)))
		heap.Push(&h, gift)
	}

	var res int64 = 0
	for i := range h {
		res += int64(h[i])
	}
	return res
}
