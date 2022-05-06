package linked_list_cycle_ii

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

func detectCycleMap(head *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})

	for head != nil {
		if _, ok := m[head]; ok {
			return head
		} else {
			m[head] = struct{}{}
		}
		head = head.Next
	}

	return nil
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast := head
	slow := head

	for fast != nil {
		if fast.Next == nil {
			return nil
		}

		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			pos := head
			for pos != slow {
				pos = pos.Next
				slow = slow.Next
			}
			return pos
		}
	}

	return nil
}
