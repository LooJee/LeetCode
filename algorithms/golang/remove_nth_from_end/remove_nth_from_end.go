package remove_nth_from_end

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	lastNode := head
	l := 1
	for lastNode.Next != nil {
		lastNode = lastNode.Next
		l++
	}

	if l == n {
		head = head.Next
		return head
	}

	i := 1
	cursor := head
	for i < l-n {
		cursor = cursor.Next
		i++
	}

	if cursor.Next != nil {
		cursor.Next = cursor.Next.Next
	}

	return head
}
