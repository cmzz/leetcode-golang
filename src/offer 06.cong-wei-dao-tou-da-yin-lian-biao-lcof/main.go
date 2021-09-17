package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	res := make([]int, 0)
	heap := make([]int, 0)

	for head != nil {
		heap = append(heap, head.Val)
		head = head.Next
	}

	for i := len(heap) - 1; i >= 0; i-- {
		res = append(res, heap[i])
	}

	return res
}

func reversePrint2(head *ListNode) []int {
	res := make([]int, 0)

	if head != nil {
		res = getVal(head, res)
	}

	return res
}

func getVal(head *ListNode, res []int) []int {
	if head == nil {
		return res
	}

	if head.Next != nil {
		res = getVal(head.Next, res)
	}

	res = append(res, head.Val)
	return res
}

func main() {
	head := &ListNode{Val: 1}
	n1 := &ListNode{Val: 2}
	n2 := &ListNode{Val: 6}
	n3 := &ListNode{Val: 4}
	n4 := &ListNode{Val: 6}
	n3.Next = n4
	n2.Next = n3
	n1.Next = n2
	head.Next = n1

	reversePrint2(head)
}
