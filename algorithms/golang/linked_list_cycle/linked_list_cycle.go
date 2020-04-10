package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println()
}

func hasCycle(head *ListNode) bool {
	quickP := head
	slowP := head

	for quickP != nil && quickP.Next != nil {
		quickP = quickP.Next.Next
		slowP = slowP.Next

		if quickP == slowP {
			return truex
		}
	}

	return false
}
