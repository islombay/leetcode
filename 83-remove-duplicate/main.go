package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var new = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}

	res := deleteDuplicates(new)
	for {
		if res != nil {
			fmt.Println(res.Val)
			res = res.Next
		}
		if res == nil {
			break

		}
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	var duplicates = make(map[int]bool)

	if head == nil {
		return head
	}

	tmp := head
	var (
		res   *ListNode = &ListNode{}
		new   *ListNode = res
		first bool      = true
	)
	for {
		if _, exists := duplicates[tmp.Val]; !exists {
			duplicates[tmp.Val] = true
			if new == nil {
				new = &ListNode{}
			}
			if !first {
				new.Next = &ListNode{}
				new = new.Next
			}
			new.Val = tmp.Val
			first = false
		}

		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}
	return res
}
