package palindromelinkedlist

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

func isPalindrome(head *ListNode) bool {
	return traverse(&head, head)
}

func traverse(head **ListNode, cur *ListNode) bool {
	if cur == nil {
		return true
	}

	state := traverse(head, cur.Next)

	if !state || (*head).Val != cur.Val {
		return false
	}

	*head = (*head).Next

	return true
}
