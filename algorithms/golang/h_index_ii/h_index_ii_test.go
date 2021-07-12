package h_index_ii

import "testing"

func TestHIndex(t *testing.T) {
	if hIndex([]int{0, 1, 3, 5, 6}) != 3 {
		t.Fatal("should be 3")
	}

	if hIndex([]int{}) != 0 {
		t.Fatal("should be 0")
	}

	if hIndex([]int{1}) != 1 {
		t.Fatal("should be 1")
	}

	if hIndex([]int{0}) != 0 {
		t.Fatal("should be 0")
	}

	if hIndex([]int{0, 1}) != 1 {
		t.Fatal("should be 1")
	}

	if hIndex([]int{0, 0, 1}) != 1 {
		t.Fatal("should be 1")
	}

	if hIndex([]int{0, 1, 1}) != 1 {
		t.Fatal("should be 1")
	}

	if hIndex([]int{100}) != 1 {
		t.Fatal("should be 1")
	}

	if hIndex([]int{11, 15}) != 2 {
		t.Fatal("should be 2")
	}
}
