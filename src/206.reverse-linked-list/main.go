package _06_reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList2(head *ListNode) *ListNode {
	var res = &ListNode{}
	var i = 0

	if head == nil {
		return nil
	}

	for head != nil {
		p := &ListNode{}
		p.Val = head.Val

		if i > 0 {
			p.Next = res
		}

		res = p

		i += 1
		head = head.Next
	}

	return res
}

// 递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
