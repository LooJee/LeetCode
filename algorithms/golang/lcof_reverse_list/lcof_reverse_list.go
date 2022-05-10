package lcof_reverse_list

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

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	t := head
	for t != nil {
		cursor := t.Next
		t.Next = prev
		prev = t
		t = cursor
	}

	return prev
}
