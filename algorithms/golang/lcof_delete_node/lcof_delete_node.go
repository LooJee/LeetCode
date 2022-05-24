package lcof_delete_node

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

func deleteNode(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}

	for t := dummyHead; t.Next != nil; t = t.Next {
		if t.Next.Val == val {
			t.Next = t.Next.Next
			break
		}
	}

	return dummyHead.Next
}
