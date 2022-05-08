package reorder_list

import "testing"

func TestReorderList(t *testing.T) {
	h := &ListNode{}
	l := h

	for i := 1; i <= 4; i++ {
		l.Next = &ListNode{Val: i}
		l = l.Next
	}

	reorderList(h.Next)

	for tmp := h.Next; tmp != nil; tmp = tmp.Next {
		t.Log(tmp.Val)
	}
}
