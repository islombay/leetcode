package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	Case1ListNode2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: -4,
			},
		},
	}
	Case1ListNode2.Next.Next.Next = Case1ListNode2

	Case2ListNode1 := &ListNode{
		Val: 1,
	}
	Case2ListNode1.Next = &ListNode{Val: 2, Next: Case2ListNode1}

	cases := []struct {
		name   string
		head   *ListNode
		result bool
	}{
		{
			name:   "1",
			result: true,
			head: &ListNode{
				Val:  3,
				Next: Case1ListNode2,
			},
		},
		{
			name:   "2",
			result: true,
			head:   Case2ListNode1,
		},
		{
			name:   "3",
			result: false,
			head:   &ListNode{Val: 1},
		},
		{
			name:   "4",
			result: false,
			head:   nil,
		},
	}

	for _, c := range cases {
		if res := hasCycle(c.head); res != c.result {
			fmt.Printf("- Test %s failed, expected %t, got %t\n", c.name, c.result, res)
		} else {
			fmt.Printf("Test %s passed, expected %t, got %t\n", c.name, c.result, res)
		}
	}
}

func hasCycle(head *ListNode) bool {
	//list := make(map[*ListNode]byte)

	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
		//if _, ok := list[head]; ok {
		//	return true
		//}
		//list[head]++
		//head = head.Next
	}
	return false
}
