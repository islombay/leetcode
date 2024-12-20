package main

import (
	"fmt"
	"reflect"
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
		want *TreeNode
	}{
		{
			id: 1,
			root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 8,
					},
					Right: &TreeNode{
						Val: 13,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 21,
					},
					Right: &TreeNode{
						Val: 34,
					},
				},
			},
			want: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 8,
					},
					Right: &TreeNode{
						Val: 13,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 21,
					},
					Right: &TreeNode{
						Val: 34,
					},
				},
			},
		},
		{
			id: 2,
			root: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
					Right: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 2,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
					Right: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 2,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
			},
			want: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
					Right: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 2,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
					Right: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 2,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		res := reverseOddLevels(c.root)
		if !reflect.DeepEqual(res, c.want) {
			fmt.Printf("Test %d failed\n", c.id)
		}
	}
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	reverse(root.Left, root.Right, 0)
	return root
}

func reverse(r1, r2 *TreeNode, level int) {
	if r1 == nil || r2 == nil {
		return
	}
	if level&1 == 0 {
		r1.Val, r2.Val = r2.Val, r1.Val
	}

	reverse(r1.Left, r2.Right, level+1)
	reverse(r1.Right, r2.Left, level+1)
}
