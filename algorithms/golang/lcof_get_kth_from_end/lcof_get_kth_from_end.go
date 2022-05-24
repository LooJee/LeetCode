package lcof_get_kth_from_end

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

func getKthFromEnd(head *ListNode, k int) *ListNode {
	l := 0
	for t := head; t != nil; t = t.Next {
		l++
	}

	if k > l {
		return nil
	}

	j := l - k

	t := head
	for i := 0; i < j; i++ {
		t = t.Next
	}

	return t
}
