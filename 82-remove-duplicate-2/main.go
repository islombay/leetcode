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
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
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

	new = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}
	res = deleteDuplicates(new)
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
	var duplicates = make(map[int]int)

	if head == nil {
		return head
	}

	var tmp = head

	for {
		duplicates[tmp.Val]++
		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}

	tmp = head

	var (
		res    *ListNode
		tmpRes      = res
		first  bool = true
	)

	for {
		if count := duplicates[tmp.Val]; count == 1 {
			if tmpRes == nil {
				res = &ListNode{}
				tmpRes = res
			}
			if !first {
				tmpRes.Next = &ListNode{}
				tmpRes = tmpRes.Next
			}
			tmpRes.Val = tmp.Val
			first = false
		}

		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}
	return res
}
