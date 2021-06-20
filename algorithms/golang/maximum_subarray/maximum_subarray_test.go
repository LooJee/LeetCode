package maximum_subarray

import "testing"

func TestMaxSubarray(t *testing.T) {
	if maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) != 6 {
		t.Fatal("should be 6")
	}

	if maxSubArray([]int{1}) != 1 {
		t.Fatal("should be 1")
	}

	if maxSubArray([]int{0}) != 0 {
		t.Fatal("should be 0")
	}

	if maxSubArray([]int{-1}) != -1 {
		t.Fatal("should be -1")
	}
}
