package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	cases := []struct {
		id   int
		nums []int
		head *ListNode
		want *ListNode
	}{
		{
			id:   1,
			nums: []int{1, 2, 3},
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
					Val:  5,
					Next: nil,
				},
			},
		},
		{
			id:   2,
			nums: []int{9, 2, 5},
			head: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 10,
					Next: &ListNode{
						Val: 9,
					},
				},
			},
			want: &ListNode{
				Val: 10,
			},
		},
	}

	for _, c := range cases {
		res := modifiedList(c.nums, c.head)
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

func modifiedList(nums []int, head *ListNode) *ListNode {
	m := make(map[int]bool)

	for _, v := range nums {
		m[v] = true
	}

	fakeHead := &ListNode{}

	prev := fakeHead
	for head != nil {
		if _, ok := m[head.Val]; ok {
			head = head.Next
			continue
		}
		prev.Next = &ListNode{
			Val: head.Val,
		}
		prev = prev.Next
		head = head.Next
	}
	return fakeHead.Next
}
