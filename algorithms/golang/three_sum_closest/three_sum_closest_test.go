package three_sum_closest

import "testing"

func TestThreeSumClosest(t *testing.T) {
	if threeSumClosest([]int{0, 0, 0}, 1) != 0 {
		t.Fatal("should be 0")
	}

	if threeSumClosest([]int{-1, 2, 1, -4}, 1) != 2 {
		t.Fatal("should be 2")
	}

	if threeSumClosest([]int{-1, 0, 1, 1, 55}, 3) != 2 {
		t.Fatal("should be 2")
	}
}
