package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func middleNode(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	slow, fast := head, head.Next

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow.Next
}
