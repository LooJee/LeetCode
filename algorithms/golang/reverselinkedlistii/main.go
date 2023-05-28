package reverselinkedlistii

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

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var (
		dummyHead = &ListNode{Next: head}
		tmp       = dummyHead
		idx       = 1
	)

	for tmp.Next != nil {
		if idx == left {
			if _, last := reverseLinkedlist(tmp.Next, idx, right); last != nil {
				tmp.Next = last
			}

			break
		}

		idx++
		tmp = tmp.Next
	}

	return dummyHead.Next
}

func reverseLinkedlist(head *ListNode, idx, end int) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}

	if idx == end {
		return head, head
	}

	idx++
	v, last := reverseLinkedlist(head.Next, idx, end)
	if v != nil {
		v.Next, head.Next = head, v.Next
	}

	return head, last
}
