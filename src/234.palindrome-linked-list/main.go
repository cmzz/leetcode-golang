package _34_palindrome_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	arr := make([]int, 0)
	if head == nil {
		return true
	}

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	var l = len(arr)
	for i := 0; i < l/2; i++ {
		j := l - i - 1

		if arr[i] != arr[j] {
			return false
		}
	}

	return true
}

func isPalindrome2(head *ListNode) bool {

	return false
}
