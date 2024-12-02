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
		val  int
		want *ListNode
	}{
		{
			id:  1,
			val: 6,
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  6,
									Next: nil,
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
		},
		{
			id:   2,
			val:  1,
			head: nil,
			want: nil,
		},
		{
			id:  3,
			val: 7,
			head: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val: 7,
						Next: &ListNode{
							Val:  7,
							Next: nil,
						},
					},
				},
			},
			want: nil,
		},
	}
	for _, c := range cases {
		res := removeElements(c.head, c.val)

		failed := false
		for c.want != nil && res != nil {
			if res.Val != c.want.Val {
				fmt.Printf("Test %d failed\n", c.id)
				failed = true
				break
			}
			res = res.Next
			c.want = c.want.Next
		}

		if failed {
			continue
		}

		if res != nil || c.want != nil {
			fmt.Printf("Test %d failed\n", c.id)
		}
	}
}

func removeElements(head *ListNode, val int) *ListNode {
	tmp := &ListNode{
		Next: nil,
	}
	vr := tmp
	for head != nil {
		if head.Val != val {
			vr.Next = head
			vr = vr.Next
		}
		head = head.Next
	}
	vr.Next = nil
	return tmp.Next
}
