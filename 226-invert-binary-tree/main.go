package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		head := stack[0]
		stack = stack[1:]

		if head.Left != nil && head.Right != nil {
			head.Left, head.Right = head.Right, head.Left
			stack = append(stack, head.Left, head.Right)
		}
	}

	return root
}
