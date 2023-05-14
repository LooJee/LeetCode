package mergeksortedlists

import "container/heap"

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

func mergeKLists(lists []*ListNode) *ListNode {
	var (
		pq = PriorityQueue{}
		l  = &ListNode{}
		p  = l
	)

	heap.Init(&pq)

	for _, ll := range lists {
		if ll == nil {
			continue
		}

		//将第一个元素插入优先队列
		heap.Push(&pq, ll)
	}

	for len(pq) > 0 {
		p.Next = heap.Pop(&pq).(*ListNode)
		//将 pop 的元素的下一个元素插入优先队列
		if p.Next.Next != nil {
			heap.Push(&pq, p.Next.Next)
		}

		p = p.Next
	}

	return l.Next
}

type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*ListNode)
	*pq = append(*pq, node)
}

// 在 pop 之前会调用 swap 将最小值和最后一个值交换位置，所以这里的 pop 是移除最后一个元素。
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}
