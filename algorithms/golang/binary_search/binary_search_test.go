package binarysearch

import "testing"

func TestSearch(t *testing.T) {
	if search([]int{-1, 0, 3, 5, 9, 12}, 9) != 4 {
		t.Fatal("should be 4")
	}

	if search([]int{-1, 0, 3, 5, 9, 12}, 2) != -1 {
		t.Fatal("should be -1")
	}

	if search([]int{-1}, -1) != 0 {
		t.Fatal("should be 0")
	}
}
