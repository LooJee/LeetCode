package add_two_numbers

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, node *ListNode

	var carry int
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + carry
		val, carry = val%10, val/10

		if head == nil {
			head = &ListNode{Val: val}
			node = head
		} else {
			node.Next = &ListNode{Val: val}
			node = node.Next
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	for ; l1 != nil; l1 = l1.Next {
		val := l1.Val + carry
		val, carry = val%10, val/10
		node.Next = &ListNode{Val: val}
		node = node.Next
	}

	for ; l2 != nil; l2 = l2.Next {
		val := l2.Val + carry
		val, carry = val%10, val/10
		node.Next = &ListNode{Val: val}
		node = node.Next
	}

	if carry != 0 {
		node.Next = &ListNode{Val: carry}
	}

	return head
}
