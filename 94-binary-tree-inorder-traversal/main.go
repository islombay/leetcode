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
		res  []int
	}{
		{
			id: 1,
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			res: []int{1, 3, 2},
		},
		{
			id:   2,
			root: nil,
			res:  []int{},
		},
		{
			id:   3,
			root: &TreeNode{Val: 1},
			res:  []int{1},
		},
	}

	for _, c := range cases {
		res := inorderTraversal(c.root)
		if len(res) != len(c.res) {
			fmt.Printf("Test %d failed, len\n", c.id)
			continue
		}
		for i := 0; i < len(res); i++ {
			if res[i] != c.res[i] {
				fmt.Printf("Test %d failed\n", c.id)
				break
			}
		}
	}
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return inorder(root)
}

func inorder(root *TreeNode) []int {
	res := []int{}
	if root.Left != nil {
		res = append(res, inorder(root.Left)...)
	}

	res = append(res, root.Val)
	if root.Right != nil {
		res = append(res, inorder(root.Right)...)
	}
	return res
}
