package reversenodesinkgroup

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	var (
		tmp = head
	)

	for i := 0; i < k; i++ {
		if tmp == nil {
			return head
		}
		tmp = tmp.Next
	}

	newHead := reverseLinkedlist(head, tmp)
	head.Next = reverseKGroup(tmp, k)

	return newHead
}

func reverseLinkedlist(head, end *ListNode) *ListNode {
	var (
		pre *ListNode
		tmp = head
	)

	for tmp != end {
		pre, tmp, tmp.Next = tmp, tmp.Next, pre
	}

	return pre
}
