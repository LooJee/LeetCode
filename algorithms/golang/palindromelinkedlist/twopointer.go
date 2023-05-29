package palindromelinkedlist

func isPalindromeO1(head *ListNode) bool {
	pre, middle := findMiddle(head)
	reversed := reverse(middle)

	for head != nil && reversed != nil {
		if head.Val != reversed.Val {
			return false
		}
		head = head.Next
		reversed = reversed.Next
	}

	pre.Next = reverse(reversed)

	return true
}

func findMiddle(head *ListNode) (*ListNode, *ListNode) {
	var (
		pre        *ListNode
		fast, slow = head, head
	)

	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		// 这种情况下节点数量为奇数，slow 指针需要向后偏移一位
		pre = slow
		slow = slow.Next
	}

	return pre, slow
}

func reverse(head *ListNode) *ListNode {
	var (
		pre *ListNode
		tmp = head
	)

	for tmp != nil {
		pre, tmp, tmp.Next = tmp, tmp.Next, pre
	}

	return pre
}
