package containerwithmostwater

import "testing"

func TestMaxArea(t *testing.T) {
	if maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) != 49 {
		t.Fatal("should be 49")
	}

	if maxArea([]int{1, 1}) != 1 {
		t.Fatal("should be 1")
	}
}
