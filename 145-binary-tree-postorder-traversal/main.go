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
		want []int
	}{
		{
			id: 1,
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val:   6,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 9,
						},
					},
				},
			},
			want: []int{4, 6, 7, 5, 2, 9, 8, 3, 1},
		},
		{
			id:   2,
			root: nil,
			want: []int{},
		},
	}

	for _, c := range cases {
		res := postorderTraversal(c.root)
		if len(res) != len(c.want) {
			fmt.Printf("Test %d failed, len\n", c.id)
			continue
		}
		for i, v := range c.want {
			if res[i] != v {
				fmt.Printf("Test %d failed\n", c.id)
				break
			}
		}
	}
}

type node struct {
	Val  int
	Next *node
}
type vector struct {
	Head *node
	Tail *node
}

func (v *vector) Prepend(val int) {
	n := &node{Val: val}
	if v.Head != nil {
		n.Next = v.Head
	}
	v.Head = n
}

func (v *vector) ToArray() []int {
	tmp := v.Head
	res := []int{}

	for tmp != nil {
		res = append(res, tmp.Val)
		tmp = tmp.Next
	}
	return res
}

func postorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	result := vector{}
	if root == nil {
		return result.ToArray()
	}

	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result.Prepend(node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return result.ToArray()
}
