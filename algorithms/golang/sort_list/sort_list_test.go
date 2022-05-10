package sort_list

import "testing"

func TestSortList(t *testing.T) {
	nums := []int{-1, 5, 3, 4, 0}

	dummyHead := &ListNode{}
	tmp := dummyHead
	for _, num := range nums {
		tmp.Next = &ListNode{Val: num}
		tmp = tmp.Next
	}

	sortList(dummyHead.Next)
}
