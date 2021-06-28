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
	p := head

	for p != nil {
		for p.Next != nil && p.Next.Val == p.Val {
			p.Next = p.Next.Next
		}
		p = p.Next
	}

	return head
}
