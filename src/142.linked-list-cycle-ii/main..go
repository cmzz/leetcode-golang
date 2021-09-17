package _42_linked_list_cycle_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针
func detectCycle(head *ListNode) *ListNode {
	var fast, slow = head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			for head != slow {
				head = head.Next
				slow = slow.Next
			}

			return slow
		}
	}

	return nil
}

// 集合
func detectCycle2(head *ListNode) *ListNode {
	c := make(map[*ListNode]struct{}, 0)

	for head != nil {
		_, ok := c[head]
		if !ok {
			c[head] = struct{}{}
		} else {
			return head
		}

		head = head.Next
	}

	return nil
}
