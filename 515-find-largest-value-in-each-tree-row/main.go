package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := []*TreeNode{root}
	stackLevel := []int{1}

	levelMax := map[int]int{1: root.Val}
	level := 1

	for len(stack) > 0 {
		root = stack[0]
		stack = stack[1:]
		level = stackLevel[0]
		stackLevel = stackLevel[1:]

		if root == nil {
			continue
		}

		val, exists := levelMax[level]

		if (exists && root.Val > val) || !exists {
			levelMax[level] = root.Val
		}

		stack = append(stack, root.Left)
		stackLevel = append(stackLevel, level+1)

		stack = append(stack, root.Right)
		stackLevel = append(stackLevel, level+1)
	}

	n := len(levelMax)
	res := make([]int, n)

	for k, v := range levelMax {
		res[k-1] = v
	}

	return res
}
