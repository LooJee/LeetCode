package lcof_min_array

import "testing"

func TestMinArray(t *testing.T) {
	if minArray([]int{3, 4, 5, 1, 2}) != 1 {
		t.Fatal("should be 1")
	}

	if minArray([]int{3, 4, 5, 6, 1, 2}) != 1 {
		t.Fatal("should be 1")
	}
	if minArray([]int{2, 2, 2, 2, 1, 2}) != 1 {
		t.Fatal("should be 1")
	}
}
