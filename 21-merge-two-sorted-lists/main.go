package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	res := mergeTwoLists(list1, list2)
	for res != nil {
		fmt.Print(res.Val, "-")
		res = res.Next
	}
	fmt.Println()
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	tmp1 := list1
	tmp2 := list2

	new := &ListNode{}
	new_tmp := new

	for {
		if tmp1 == nil {
			new_tmp.Val = tmp2.Val
			new_tmp.Next = tmp2.Next
			break
		} else if tmp2 == nil {
			new_tmp.Val = tmp1.Val
			new_tmp.Next = tmp1.Next
			break
		}
		if tmp1.Val > tmp2.Val {
			new_tmp.Val = tmp2.Val
			tmp2 = tmp2.Next
		} else {
			new_tmp.Val = tmp1.Val
			tmp1 = tmp1.Next
		}
		new_tmp.Next = &ListNode{}
		new_tmp = new_tmp.Next
	}
	return new
}
