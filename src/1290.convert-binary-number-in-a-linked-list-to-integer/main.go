package _290_convert_binary_number_in_a_linked_list_to_integer

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	res := 0

	for head != nil {
		res = res*2 + head.Val
		head = head.Next
	}

	return res
}
