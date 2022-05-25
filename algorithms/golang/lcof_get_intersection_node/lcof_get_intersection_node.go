package lcof_get_intersection_node

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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1, l2 := headA, headB

	for l1 != nil && l2 != nil {
		l1 = l1.Next
		l2 = l2.Next
	}

	l3, l4 := headA, headB
	if l1 != nil {
		for l1 != nil && l3 != nil {
			l1 = l1.Next
			l3 = l3.Next
		}
	}

	if l2 != nil {
		for l2 != nil && l4 != nil {
			l2 = l2.Next
			l4 = l4.Next
		}
	}

	for l3 != nil && l4 != nil {
		if l3 == l4 {
			return l3
		}
		l3 = l3.Next
		l4 = l4.Next
	}

	return nil
}
