package remove_duplicates_from_sorted_list

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var p, q = head, head

	for p != nil {
		if p.Val != q.Val {
			q.Next = p
			q = p
		}
		p = p.Next
	}

	q.Next = nil

	return head
}
