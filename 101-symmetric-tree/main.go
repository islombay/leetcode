package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	cases := []struct {
		id   int
		root *TreeNode
		want bool
	}{
		{
			id: 1,
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: true,
		},
		{
			id: 2,
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: false,
		},
	}

	for _, c := range cases {
		if res := isSymmetric(c.root); res != c.want {
			fmt.Printf("Test %d failed, expected %t, got %t\n", c.id, c.want, res)
		}
	}
}

func isSymmetric(root *TreeNode) bool {
	left, right := root.Left, root.Right
	stackLeft, stackRight := []*TreeNode{}, []*TreeNode{}

	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	}

	stackLeft = append(stackLeft, left)
	stackRight = append(stackRight, right)

	for len(stackLeft) > 0 && len(stackRight) > 0 {
		left, right = stackLeft[len(stackLeft)-1], stackRight[len(stackRight)-1]
		stackLeft, stackRight = stackLeft[:len(stackLeft)-1], stackRight[:len(stackRight)-1]

		if left.Val != right.Val {
			return false
		}

		if left.Left != nil && right.Right != nil {
			stackLeft = append(stackLeft, left.Left)
			stackRight = append(stackRight, right.Right)
		} else if left.Left == nil && right.Right == nil {

		} else if left.Left == nil || right.Right == nil {
			return false
		}

		if left.Right != nil && right.Left != nil {
			stackLeft = append(stackLeft, left.Right)
			stackRight = append(stackRight, right.Left)
		} else if left.Right == nil && right.Left == nil {

		} else if left.Right == nil || right.Left == nil {
			return false
		}
	}

	return true
}
