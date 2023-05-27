package reverselinkedlist

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
	if head == nil {
		return nil
	}

	var (
		oldDummy, newDummy = &ListNode{Next: head}, &ListNode{Next: nil}
	)

	for oldDummy.Next != nil {
		node := oldDummy.Next
		oldDummy.Next = node.Next
		node.Next = newDummy.Next
		newDummy.Next = node
	}

	return newDummy.Next
}
