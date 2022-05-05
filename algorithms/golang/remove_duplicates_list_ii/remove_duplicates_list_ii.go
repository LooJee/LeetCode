package remove_duplicates_list_ii

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	newHead := &ListNode{
		Val:  0,
		Next: head,
	}

	p := newHead
	for p.Next != nil && p.Next.Next != nil {
		if p.Next.Val == p.Next.Next.Val {
			v := p.Next.Val
			p.Next = p.Next.Next
			for p.Next != nil && p.Next.Val == v {
				p.Next = p.Next.Next
			}
		} else {
			p = p.Next
		}
	}

	return newHead.Next
}
