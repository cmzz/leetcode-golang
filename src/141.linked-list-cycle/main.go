package _41_linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 * https://leetcode-cn.com/problems/linked-list-cycle/
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	var slow, fast = head.Next, head.Next.Next

	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}

/**
 * 遍历
 */
func hasCycle2(head *ListNode) bool {
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return true
		}

		seen[head] = struct{}{}
		head = head.Next
	}

	return false
}
