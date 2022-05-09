package peak_index_in_mountain_array

import "testing"

func TestPeakIndexInMountainArray(t *testing.T) {
	if peakIndexInMountainArray([]int{3, 5, 3, 2, 0}) != 1 {
		t.Fatal("should be 1")
	}
}
