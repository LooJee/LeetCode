package searchrange

import "testing"

func TestSearchRange(t *testing.T) {
	t.Log(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}
