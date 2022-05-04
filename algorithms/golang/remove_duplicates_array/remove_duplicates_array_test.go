package remove_duplicates_array

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	l := removeDuplicates(nums)
	t.Log(l, nums[:l])
	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	l = removeDuplicates(nums)
	t.Log(l, nums[:l])
}
