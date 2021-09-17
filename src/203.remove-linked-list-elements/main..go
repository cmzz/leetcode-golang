package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	p := &ListNode{}
	p.Next = head
	res := p

	for p != nil && p.Next != nil {
		if p.Next.Val == val {
			if p.Next.Next != nil {
				p.Next = p.Next.Next
			} else {
				p.Next = nil
			}
		} else {
			p = p.Next
		}
	}

	return res.Next
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

	res := removeElements(head, 6)
	fmt.Printf("\n %#v", res)
}
