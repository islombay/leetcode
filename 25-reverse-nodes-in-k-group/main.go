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
		k    int
		res  *ListNode
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
			res: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
		},
		{
			id: 2,
			k:  3,
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
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
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
		},
	}

	for _, c := range cases {
		head := reverseKGroup(c.head, c.k)
		tmpHead := c.res
		failed := false

		for tmpHead != nil && head != nil {
			if tmpHead.Val != head.Val {
				failed = true
				fmt.Printf("Test %d failed\n", c.id)
				break
			}
			tmpHead = tmpHead.Next
			head = head.Next
		}
		if !failed && head != nil {
			fmt.Printf("Test %d failed\n", c.id)
			continue
		}

		if !failed && tmpHead != nil {
			fmt.Printf("Test %d failed\n", c.id)
		}
	}
	fmt.Println("Finish")
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	arr := []int{}

	tmpHead := &ListNode{}
	tmpNode := tmpHead

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	for i := k - 1; i < len(arr); i += k {
		for j := i; j >= i-k+1; j-- {
			tmpNode.Next = &ListNode{
				Val: arr[j],
			}
			tmpNode = tmpNode.Next
		}
	}
	for i := len(arr) - (len(arr) % k); i < len(arr); i++ {
		tmpNode.Next = &ListNode{
			Val: arr[i],
		}
		tmpNode = tmpNode.Next
	}
	return tmpHead.Next
}
