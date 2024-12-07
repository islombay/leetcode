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
		p    *TreeNode
		q    *TreeNode
		want bool
	}{
		{
			id: 1,
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			q: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			want: true,
		},
		{
			id: 2,
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			q: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			want: false,
		},
	}

	for _, c := range cases {
		if res := isSameTree(c.p, c.q); res != c.want {
			fmt.Printf("Test %d failed, expected %t, got %t\n", c.id, c.want, res)
		}
	}
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	stackP := []*TreeNode{}
	stackQ := []*TreeNode{}

	stackP = append(stackP, p)
	stackQ = append(stackQ, q)

	for len(stackP) > 0 && len(stackQ) > 0 {
		tmpP := stackP[len(stackP)-1]
		tmpQ := stackQ[len(stackQ)-1]

		stackP = stackP[:len(stackP)-1]
		stackQ = stackQ[:len(stackQ)-1]

		if tmpP.Val != tmpQ.Val {
			return false
		}

		if tmpP.Left != nil && tmpQ.Left != nil {
			stackP = append(stackP, tmpP.Left)
			stackQ = append(stackQ, tmpQ.Left)
		} else if tmpP.Left != nil || tmpQ.Left != nil {
			return false
		}

		if tmpP.Right != nil && tmpQ.Right != nil {
			stackP = append(stackP, tmpP.Right)
			stackQ = append(stackQ, tmpQ.Right)
		} else if tmpP.Right != nil || tmpQ.Right != nil {
			return false
		}
	}
	return true
}
