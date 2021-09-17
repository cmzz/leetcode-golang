package _76_middle_of_the_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	res := make([]*ListNode, 0)
	for head != nil {
		res = append(res, head)
		head = head.Next
	}

	return res[len(res)/2]
}

func middleNode2(head *ListNode) *ListNode {
	var count = 0
	var p = head
	for head != nil {
		count += 1
		head = head.Next
	}

	var i = 0
	for p != nil {
		if i == count/2 {
			return p
		}

		p = p.Next
		i += 1
	}

	return nil
}

func middleNode3(head *ListNode) *ListNode {
	var fast = head

	for fast != nil {
		if fast.Next != nil {
			head = head.Next

			if fast.Next.Next != nil {
				fast = fast.Next.Next
			} else {
				break
			}
		} else {
			break
		}
	}

	return head
}
