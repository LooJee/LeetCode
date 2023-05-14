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

// 使用快慢指针可以不使用 O(1) 的内存空间解决问题
// 思路是，fast 通过 2k 步，而 slow 使用 k 步，最终两个指针相遇。
// 假设，环入口到相遇点的距离是 m ，那么从 head 出发，经过 k-m 步可以到达环入口，同样的，从环内相遇点出发，经过 k-m 步可以到达环的入口位置
// 由此推断，当两个指针在环内相遇后，一个指针回到 head，然后两个指针同时单步前行，它们会在环入口相遇
func detectCycle(head *ListNode) *ListNode {
	var slow, fast = head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	fast = head

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}
