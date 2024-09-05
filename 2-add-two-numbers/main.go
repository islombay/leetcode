package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	cases := []struct {
		name   string
		l1     *ListNode
		l2     *ListNode
		result *ListNode
	}{
		{
			name: "1",
			l1: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
			l2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			result: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 8,
					},
				},
			},
		},
		{
			name: "2",
			l1: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
								Next: &ListNode{
									Val: 9,
									Next: &ListNode{
										Val: 9,
									},
								},
							},
						},
					},
				},
			},
			l2: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
						},
					},
				},
			},
			result: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val: 0,
									Next: &ListNode{
										Val: 0,
										Next: &ListNode{
											Val: 1,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "3",
			l1: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 9,
					},
				},
			},
			l2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 9,
						},
					},
				},
			},
			result: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 1,
							},
						},
					},
				},
			},
		},
		{
			name:   "4",
			l1:     &ListNode{Val: 5},
			l2:     &ListNode{Val: 5},
			result: &ListNode{Val: 0, Next: &ListNode{Val: 1}},
		},
	}

	for _, c := range cases {
		res := addTwoNumbers(c.l1, c.l2)
		for res != nil {
			if res.Val != c.result.Val {
				fmt.Printf("- Test %s failed, expected %v, got %v\n", c.name, c.result, res)
				break
			}
			res = res.Next
			c.result = c.result.Next
		}
		fmt.Printf("Test %s passed\n", c.name)
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var (
		additional = 0
		head       = l1
		last       *ListNode
	)

	for l1 != nil && l2 != nil {
		l1.Val = l1.Val + l2.Val + additional
		additional = l1.Val / 10
		l1.Val = l1.Val % 10

		if l1.Next == nil {
			l1.Next = l2.Next
			last = l1
			l1 = l1.Next
			break
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		if additional != 0 {
			l1.Val += additional
			additional = l1.Val / 10
			l1.Val = l1.Val % 10

			if l1.Next == nil {
				last = l1
			}

			l1 = l1.Next

		} else {
			break
		}
	}

	if additional != 0 {
		last.Next = &ListNode{
			Val: additional,
		}
	}
	return head
}
