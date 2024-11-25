package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ha, hb := headA, headB
	visit := make(map[*ListNode]int)

	for ha != nil && hb != nil {
		visit[ha]++
		if visit[ha] >= 2 {
			return ha
		}

		visit[hb]++
		if visit[hb] >= 2 {
			return hb
		}

		ha = ha.Next
		hb = hb.Next
	}

	for ha != nil {
		visit[ha]++
		if visit[ha] >= 2 {
			return ha
		}
		ha = ha.Next
	}

	for hb != nil {
		visit[hb]++
		if visit[hb] >= 2 {
			return hb
		}
		hb = hb.Next
	}
	return nil
}
