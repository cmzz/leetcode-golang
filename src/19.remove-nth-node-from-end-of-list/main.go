package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res := &ListNode{}
	res.Next = head
	var cur = res
	var pre = &ListNode{}

	i := 1
	for head != nil {
		if i < n {
			head = head.Next
		} else {
			head = head.Next
			pre = cur
			cur = cur.Next
		}
		i++
	}

	pre.Next = pre.Next.Next
	return res.Next
}

func main() {
	head := &ListNode{Val: 1, Next: nil}
	head.Next = &ListNode{Val: 2, Next: nil}
	res := removeNthFromEnd(head, 2)

	fmt.Println("\n =====")
	for res != nil {
		fmt.Printf("\n %d", res.Val)
		res = res.Next
	}
}
