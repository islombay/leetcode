package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	cases := []struct {
		id   int
		head *ListNode
		want int
	}{
		{
			id: 1,
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
			want: 5,
		},
		{
			id:   2,
			want: 18880,
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val: 0,
									Next: &ListNode{
										Val: 1,
										Next: &ListNode{
											Val: 1,
											Next: &ListNode{
												Val: 1,
												Next: &ListNode{
													Val: 0,
													Next: &ListNode{
														Val: 0,
														Next: &ListNode{
															Val: 0,
															Next: &ListNode{
																Val: 0,
																Next: &ListNode{
																	Val: 0,
																	Next: &ListNode{
																		Val:  0,
																		Next: nil,
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		if res := getDecimalValue(c.head); res != c.want {
			fmt.Printf("Test %d failed, expected %d, got %d\n", c.id, c.want, res)
		}
	}
}

func getDecimalValue(head *ListNode) int {
	s := strings.Builder{}
	for head != nil {
		s.WriteString(strconv.Itoa(head.Val))
		head = head.Next
	}
	f := s.String()
	res, _ := strconv.ParseInt(f, 2, 64)
	return int(res)
}
