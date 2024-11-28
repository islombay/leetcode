package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	order []*TreeNode
	index int
}

func (s *Stack) Push(node *TreeNode) {
	if len(s.order) <= s.index {
		s.order = append(s.order, node)
	} else {
		s.order[s.index] = node
	}
	s.index++
}

func (s *Stack) Pop() *TreeNode {
	defer func() {
		if s.index != 0 {
			s.index--
		}
	}()
	return s.Top()
}

func (s *Stack) Top() *TreeNode {
	if s.index <= 0 {
		return nil
	}
	return s.order[s.index-1]
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
			res: []int{1, 2, 3},
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
		{
			id: 4,
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 4},
					Right: &TreeNode{
						Val:   5,
						Left:  &TreeNode{Val: 6},
						Right: &TreeNode{Val: 7},
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val:  8,
						Left: &TreeNode{Val: 9},
					},
				},
			},
			res: []int{1, 2, 4, 5, 6, 7, 3, 8, 9},
		},
	}

	for _, c := range cases {
		res := preorderTraversal2(c.root)
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

func preorderTraversal2(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, root.Val)

	if root.Left != nil {
		res = append(res, preorderTraversal2(root.Left)...)
	}

	if root.Right != nil {
		res = append(res, preorderTraversal2(root.Right)...)
	}
	return res
}

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	stack := Stack{
		order: []*TreeNode{},
		index: 0,
	}
	tmp := root

	for tmp != nil {
		res = append(res, tmp.Val)
		stack.Push(tmp)

		for tmp.Left != nil {
			tmp = tmp.Left
			stack.Push(tmp)
			res = append(res, tmp.Val)
		}
		for stack.Top() != nil {
			tmp = stack.Pop()
			for tmp.Right != nil {
				tmp = tmp.Right
				stack.Pop()
				stack.Push(tmp)
				res = append(res, tmp.Val)

				for tmp.Left != nil {
					tmp = tmp.Left
					stack.Push(tmp)
					res = append(res, tmp.Val)
				}
			}
		}
		tmp = stack.Pop()
	}

	return res
}
