package rotate_list

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

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	l := 1
	last := head
	for last.Next != nil {
		l++
		last = last.Next
	}

	pos := k % l
	if pos == 0 {
		return head
	}

	i := 1
	t := head
	for i < (l - pos) {
		i++
		t = t.Next
	}

	last.Next = head
	head = t.Next
	t.Next = nil

	return head
}
