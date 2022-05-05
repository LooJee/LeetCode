package partition_list

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	smaller := &ListNode{}
	sp := smaller
	bigger := &ListNode{}
	bp := bigger

	for t := head; t != nil; t = t.Next {
		if t.Val < x {
			sp.Next = t
			sp = sp.Next
		} else {
			bp.Next = t
			bp = bp.Next
		}
	}
	bp.Next = nil
	sp.Next = bigger.Next

	return smaller.Next
}
