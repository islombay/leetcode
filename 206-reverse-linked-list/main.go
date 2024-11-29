package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	cases := []struct {
		id   int
		head *ListNode
		res  *ListNode
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
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			res: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  1,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		res := reverseList(c.head)
		for res != nil && res.Next != nil {
			if res.Val != c.res.Val {
				fmt.Printf("Test %d failed\n", c.id)
				break
			}
			res = res.Next
			c.res = c.res.Next
			if res == nil && c.res != nil {
				fmt.Printf("Test %d failed\n", c.id)
				break
			} else if c.res == nil && res != nil {
				fmt.Printf("Test %d failed\n", c.id)
				break
			}
		}
	}
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	var last *ListNode = nil
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		last = head
		head = next
	}
	return last
}
