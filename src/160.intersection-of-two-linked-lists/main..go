package _60_intersection_of_two_linked_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var a, b = headA, headB
	if a == nil || b == nil {
		return nil
	}

	for a != b {
		a = a.Next
		b = b.Next

		if a == nil && b == nil {
			return nil
		}

		if a == nil {
			a = headB
		}

		if b == nil {
			b = headA
		}
	}

	return a
}
