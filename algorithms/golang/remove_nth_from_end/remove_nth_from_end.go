package remove_nth_from_end

type ListNode struct {
	Val  int
	Next *ListNode
}

// 需要删除倒数第 n 个节点，实际是要找到倒数第 n+1 个节点，然后将其的 Next 指向倒数第 n-1 个节点。
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var dummyHead = &ListNode{Next: head}
	var x, y = head, dummyHead

	// 新建一个指针 x 从 head 开始找到第 n+1 个元素，此时剩余的元素个数是 len - n
	for i := 1; i < n && x != nil; i++ {
		x = x.Next
	}

	if x == nil {
		return head
	}

	//一直找到最后一个元素
	for x.Next != nil {
		x = x.Next
		y = y.Next
	}

	y.Next = y.Next.Next

	return dummyHead.Next
}
