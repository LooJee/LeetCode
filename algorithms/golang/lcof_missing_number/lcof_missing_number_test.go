package lcof_missing_number

import "testing"

func TestMissingNumber(t *testing.T) {
	if missingNumber([]int{0, 1, 3}) != 2 {
		t.Fatal("should be 2")
	}

	if missingNumber([]int{1, 2, 3}) != 0 {
		t.Fatal("should be 0")
	}
}
