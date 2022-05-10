package sort_list

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

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	fast := head
	slow := head

	if fast != tail {
		fast = fast.Next
		slow = slow.Next

		if fast.Next != tail {
			fast = fast.Next
		}
	}

	return merge(sort(head, slow), sort(slow, tail))
}

func merge(l, r *ListNode) *ListNode {
	ll := &ListNode{}
	tmp := ll
	for l != nil && r != nil {
		if l.Val <= r.Val {
			tmp.Next = l
			l = l.Next
		} else {
			tmp.Next = r
			r = r.Next
		}
		tmp = tmp.Next
	}

	if l != nil {
		tmp.Next = l
	}

	if r != nil {
		tmp.Next = r
	}

	return ll.Next
}
