package partition_list

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	var (
		left, right       = &ListNode{}, &ListNode{}
		leftTmp, rightTmp = left, right
	)

	for head != nil {
		if head.Val < x {
			leftTmp.Next = head
			leftTmp = leftTmp.Next
		} else {
			rightTmp.Next = head
			rightTmp = right.Next
		}
		head = head.Next
	}

	leftTmp.Next = right.Next
	rightTmp.Next = nil

	return left.Next
}
