package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	if head != nil && head.Next != nil && head.Next.Next == nil {
		return head.Next.Val == head.Val
	}

	fast, slow := head, head
	count := 1
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		count += 2
	}

	if fast.Next != nil {
		count++
	}

	slow = slow.Next
	var prev *ListNode
	tmp := slow
	for tmp != nil {
		next := tmp.Next
		tmp.Next = prev
		prev = tmp
		tmp = next
	}

	for head != nil && prev != nil {
		if head.Val != prev.Val {
			return false
		}
		head = head.Next
		prev = prev.Next
	}
	return true
}
