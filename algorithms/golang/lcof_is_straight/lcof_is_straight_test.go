package lcof_is_straight

import "testing"

func TestIsStraight(t *testing.T) {
	t.Log(isStraight([]int{1, 2, 3, 4, 5}))
	t.Log(isStraight([]int{1, 2, 2, 4, 5}))
	t.Log(isStraight([]int{0, 0, 1, 4, 5}))
	t.Log(isStraight([]int{5, 4, 1, 0, 0}))
}
