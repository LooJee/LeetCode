package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	quickP := head
	slowP := head

	for quickP != nil && quickP.Next != nil {
		quickP = quickP.Next.Next
		slowP = slowP.Next

		if quickP == slowP {
			return true
		}
	}

	return false
}
