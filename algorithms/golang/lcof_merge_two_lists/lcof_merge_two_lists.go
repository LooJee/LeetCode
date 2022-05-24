package lcof_merge_two_lists

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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}

	t := dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			t.Next = l1
			l1 = l1.Next
		} else {
			t.Next = l2
			l2 = l2.Next
		}
		t = t.Next
	}

	if l1 != nil {
		t.Next = l1
	}

	if l2 != nil {
		t.Next = l2
	}

	return dummyHead.Next
}
