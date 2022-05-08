package reorder_list

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

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	r := slow.Next
	slow.Next = nil
	var prev *ListNode
	for r != nil {
		cursor := r.Next
		r.Next = prev
		prev = r
		r = cursor
	}

	r = prev
	l := head
	for l != nil && r != nil {
		rt := r.Next
		lt := l.Next

		l.Next = r
		l = lt

		r.Next = l
		r = rt
	}
}
