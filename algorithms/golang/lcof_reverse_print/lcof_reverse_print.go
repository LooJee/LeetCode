package lcof_reverse_print

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

func reversePrint(head *ListNode) []int {
	l := 0
	for t := head; t != nil; t = t.Next {
		l++
	}

	a := make([]int, l)

	for t := head; t != nil; t = t.Next {
		a[l-1] = t.Val
		l--
	}

	return a
}
