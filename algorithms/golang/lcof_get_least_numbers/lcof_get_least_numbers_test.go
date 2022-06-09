package lcof_get_least_numbers

import "testing"

func TestGetLeastNumbers(t *testing.T) {
	t.Log(getLeastNumbers([]int{3, 2, 1}, 2))
	t.Log(getLeastNumbers([]int{0, 1, 2, 1}, 1))
}
