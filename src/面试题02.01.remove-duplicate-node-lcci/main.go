package 面试题02_01_remove_duplicate_node_lcci

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	m := make(map[int]int, 0)
	h := head

	if head == nil {
		return head
	}

	var p = &ListNode{}
	for head != nil {
		_, ok := m[head.Val]
		if !ok {
			m[head.Val] = head.Val
			p = head
		} else {
			p.Next = head.Next
		}

		head = head.Next
	}

	return h
}
