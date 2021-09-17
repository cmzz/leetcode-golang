package _2_02_kth_node_from_end_of_list_lcci

type ListNode struct {
	Val  int
	Next *ListNode
}

// 先统计连标长度
// 到数k个元素位置为 len - k
func kthToLast(head *ListNode, k int) int {
	var h = head
	if h == nil {
		return 0
	}

	var len = 0
	for h != nil {
		len += 1
		h = h.Next
	}

	for j := 0; j <= len-k; j++ {
		if j == len-k {
			return head.Val
		} else {
			head = head.Next
		}
	}

	return 0
}

// 2个指针，开始都指向投
// 其中一个后移 K-1个节点
// 同步后移2个指针，直至后移k节点的指针到达为节点
// 此时另一节点即为倒数k个节点
func kthToLast2(head *ListNode, k int) int {
	var p = head
	for i := 0; i <= k-1; i++ {
		p = p.Next
	}

	for p != nil {
		p = p.Next
		head = head.Next
	}

	return head.Val
}

// 递归
func kthToLast3(head *ListNode, k int) int {
	_, v := getVal(head, k)
	return v
}

func getVal(head *ListNode, k int) (int, int) {
	var i = 1
	var v = 0

	if head.Next != nil {
		i, v = getVal(head.Next, k)
		i += 1
	}

	if i == k {
		return i, head.Val
	}

	return i, v
}
