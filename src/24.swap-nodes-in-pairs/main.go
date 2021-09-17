package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	res := &ListNode{}
	res.Next = head
	cur := res

	for cur.Next != nil && cur.Next.Next != nil {
		n1, n2, n3, n4 := cur, cur.Next, cur.Next.Next, cur.Next.Next.Next
		n1.Next = n3
		n3.Next = n2
		n2.Next = n4

		cur = cur.Next.Next
	}

	return res.Next
}
