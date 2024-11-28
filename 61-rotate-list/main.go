package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	cases := []struct {
		id   int
		head *ListNode
		k    int
		want *ListNode
	}{
		{
			id: 1,
			k:  2,
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  3,
								Next: nil,
							},
						},
					},
				},
			},
		},
		{
			id:   2,
			k:    0,
			head: &ListNode{},
			want: &ListNode{},
		},
		{
			id:   3,
			k:    1,
			head: &ListNode{Val: 1},
			want: &ListNode{Val: 1},
		},
		{
			id: 4,
			k:  2,
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
		},
	}

	for _, c := range cases {
		res := rotateRight(c.head, c.k)
		failed := false
		for res != nil && c.want != nil {
			if res.Val != c.want.Val {
				failed = true
				fmt.Printf("Test %d failed\n", c.id)
				break
			}
			res = res.Next
			c.want = c.want.Next
		}
		if !failed && res != nil {
			fmt.Printf("Test %d failed\n", c.id)
			continue
		}

		if !failed && c.want != nil {
			fmt.Printf("Test %d failed\n", c.id)
		}
	}
}

func Size(o *ListNode) int {
	tmp := o
	count := 0
	for tmp != nil {
		count++
		tmp = tmp.Next
	}
	return count
}

func rotateRight(head *ListNode, k int) *ListNode {
	first := head
	var newFirst *ListNode

	size := Size(head)
	if size == 0 || k == 0 || size == 1 {
		return head
	}
	steps := size - (k % size)
	if steps-size == 0 {
		return head
	}
	tmp := head

	for steps != 1 {
		tmp = tmp.Next
		steps--
	}

	newFirst = tmp.Next
	tmp.Next = nil

	tmp = newFirst
	for tmp != nil {
		if tmp.Next == nil {
			tmp.Next = first
			break
		}
		tmp = tmp.Next
	}
	return newFirst
}
