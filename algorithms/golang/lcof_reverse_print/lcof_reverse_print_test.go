package lcof_reverse_print

import "testing"

func TestReversePrint(t *testing.T) {
	a := []int{1, 3, 2}
	list := &ListNode{}

	tmp := list
	for i := 0; i < len(a); i++ {
		tmp.Next = &ListNode{Val: a[i]}
		tmp = tmp.Next
	}

	t.Log(reversePrint(list.Next))
}
