package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	cases := []struct {
		id   int
		root *TreeNode
		want []string
	}{
		{
			id: 1,
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			want: []string{"1->2->5", "1->3"},
		},
	}

	for _, c := range cases {
		res := binaryTreePaths(c.root)
		if !reflect.DeepEqual(res, c.want) {
			fmt.Printf("Test %d failed\n", c.id)
		}
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	node       *TreeNode
	anchestors []int
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	intResults := [][]int{}
	stack := []*Node{
		{
			node:       root,
			anchestors: []int{},
		},
	}

	var head = stack[0]

	for len(stack) > 0 {
		head = stack[0]
		stack = stack[1:]

		newAnchestor := append([]int{}, head.anchestors...)
		newAnchestor = append(newAnchestor, head.node.Val)

		if head.node.Left != nil {
			stack = append(stack, &Node{
				node:       head.node.Left,
				anchestors: newAnchestor,
			})
		}
		if head.node.Right != nil {
			stack = append(stack, &Node{
				node:       head.node.Right,
				anchestors: newAnchestor,
			})
		}

		if head.node.Left == nil && head.node.Right == nil {
			intResults = append(intResults, newAnchestor)
		}
	}
	res := []string{}
	for _, nums := range intResults {
		tmp := strings.Builder{}
		for i, num := range nums {
			tmp.WriteString(strconv.Itoa(num))
			if i != len(nums)-1 {
				tmp.WriteString("->")
			}
		}
		res = append(res, tmp.String())
	}
	return res
}
