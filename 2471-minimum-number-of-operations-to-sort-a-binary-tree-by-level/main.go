package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	cases := []struct {
		id   int
		root *TreeNode
		want int
	}{
		{
			id: 1,
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 9,
						},
					},
					Right: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 10,
						},
					},
				},
			},
			want: 3,
		},
		{
			id: 2,
			root: &TreeNode{
				Val: 11,
				Right: &TreeNode{
					Val: 43,
					Right: &TreeNode{
						Val: 17,
					},
				},
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 27,
					},
					Right: &TreeNode{
						Val: 29,
						Left: &TreeNode{
							Val: 18,
						},
					},
				},
			},
			want: 2,
		},
	}

	for _, c := range cases {
		if res := minimumOperations(c.root); res != c.want {
			fmt.Printf("Test %d failed, expected %d got %d\n", c.id, c.want, res)
		}
	}
}

func minimumOperations(root *TreeNode) int {
	levelArray := map[int][]int{}
	stack := []*TreeNode{root}
	stackIdx := []int{1}

	for len(stack) > 0 {
		root = stack[0]
		rootIdx := stackIdx[0]
		stack = stack[1:]
		stackIdx = stackIdx[1:]

		if root == nil {
			continue
		}

		if rootIdx != 1 {
			_, ok := levelArray[rootIdx]
			if !ok {
				levelArray[rootIdx] = []int{}
			}
			levelArray[rootIdx] = append(levelArray[rootIdx], root.Val)
		}

		stack = append(stack, root.Left)
		stackIdx = append(stackIdx, rootIdx+1)

		stack = append(stack, root.Right)
		stackIdx = append(stackIdx, rootIdx+1)
	}

	counter := 0

	for _, v := range levelArray {
		//for i := 1; i < len(v); i++ {
		//	if v[i] < v[i-1] {
		//		counter++
		//		i++
		//	}
		//}
		counter += countSwaps(v)
	}
	return counter
}

func countSwaps(arr []int) int {
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)

	indexMap := make(map[int]int)
	for i, val := range sorted {
		indexMap[val] = i
	}

	visited := make([]bool, len(arr))
	swaps := 0

	for i := 0; i < len(arr); i++ {
		if visited[i] || arr[i] == sorted[i] {
			continue
		}

		cycleSize := 0
		j := i

		for !visited[j] {
			visited[j] = true
			j = indexMap[arr[j]]
			cycleSize++
		}

		if cycleSize > 0 {
			swaps += cycleSize - 1
		}
	}

	return swaps
}
