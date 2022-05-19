package find_kth_positive

import "testing"

func TestFindKthPositive(t *testing.T) {
	t.Log(findKthPositive([]int{2, 3, 4, 7, 11}, 5))
	t.Log(findKthPositive([]int{1, 2, 3, 4, 5}, 7))
	t.Log(findKthPositive([]int{1, 2, 3, 4}, 2))
	t.Log(findKthPositive([]int{8, 11, 16, 20, 29, 30, 32, 33, 37, 39, 42, 44, 46, 47, 48, 50, 52, 56, 60, 63, 64, 65, 68, 70, 72, 74, 80}, 45))
}
