package main

import "fmt"

/*
给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。
你可以假设除了数字 0 之外，这两个数字都不会以零开头。

示例：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := new(ListNode)
	current := res

	a := l1
	b := l2

	carry := 0
	x := 0
	y := 0
	sum := 0

	for a != nil || b != nil {
		if a == nil {
			x = 0
		} else {
			x = a.Val
		}

		if b == nil {
			y = 0
		} else {
			y = b.Val
		}

		sum = x + y + carry
		carry = sum / 10

		current.Val = sum % 10

		if a != nil {
			a = a.Next
		}

		if b != nil {
			b = b.Next
		}

		if a != nil || b != nil {
			current.Next = new(ListNode)
			current = current.Next
		}
	}

	if carry > 0 {
		current.Next = &ListNode{Val:carry, Next:nil}
	}

	return res
}

/**
 * 创建列表
 */
func makeList(nums []int) *ListNode  {
	l := new(ListNode)
	curr := l

	for i, v := range nums {
		curr.Val = v

		if	i < len(nums) - 1 {
			curr.Next = new(ListNode)
			curr = curr.Next
		}
	}

	return l
}

/**
 * 打印列表
 */
func printList(l *ListNode)  {
	var t []int

	for l != nil {
		t = append(t, l.Val)
		l = l.Next
	}

	fmt.Printf("%v\n", t)
}

func main()  {
	l1 := makeList([]int{1,8,3})
	l2 := makeList([]int{0})
	l3 := addTwoNumbers(l1, l2)

	printList(l1)
	printList(l2)
	printList(l3)

	
}
