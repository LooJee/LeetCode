package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	quick, slow := head, head

	for quick != nil && quick.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next

		if quick == slow {
			return true
		}
	}

	return false
}
