package main

import "fmt"

func main() {
	cases := []struct {
		id   int
		head *ListNode
		want *ListNode
	}{
		{
			id: 1,
			head: &ListNode{
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
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
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
		{
			id: 2,
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
				Val: 1,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  3,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	for _, c := range cases {
		reorderList(c.head)

		for c.head != nil && c.want != nil {
			if c.head.Val != c.want.Val {
				fmt.Printf("Test %d failed\n", c.id)
				break
			}
			c.head = c.head.Next
			c.want = c.want.Next
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	stack := []*ListNode{}

	tmpHead := head
	for tmpHead != nil {
		stack = append(stack, tmpHead)
		tmpHead = tmpHead.Next
	}

	tmpHead = head

	for tmpHead != nil && len(stack) != 0 {
		tmpNext := tmpHead.Next
		tail := stack[len(stack)-1]
		if tmpHead == tail || tmpNext == tail {
			tail.Next = nil
			break
		}
		stack = stack[1 : len(stack)-1]
		tail.Next = tmpNext
		tmpHead.Next = tail
		tmpHead = tmpNext
	}
}
