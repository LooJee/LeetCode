package intersectionoftwolinkedlists

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

// 分别得到两点的长度，将较长的链表先移动，直到两个链表剩余的长度相同，然后同时移动两个链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var (
		tempA, tempB = headA, headB
		lenA, lenB   = 0, 0

		longger, shorter       *ListNode
		longgerLen, shorterLen int
	)

	for tempA != nil {
		tempA = tempA.Next
		lenA++
	}

	for tempB != nil {
		tempB = tempB.Next
		lenB++
	}

	if lenA > lenB {
		longger, shorter = headA, headB
		longgerLen, shorterLen = lenA, lenB
	} else {
		longger, shorter = headB, headA
		longgerLen, shorterLen = lenB, lenA
	}

	for diffLen := longgerLen - shorterLen; diffLen > 0; diffLen-- {
		longger = longger.Next
	}

	for longger != nil {
		if longger == shorter {
			return longger
		}

		longger = longger.Next
		shorter = shorter.Next
	}

	return nil
}
