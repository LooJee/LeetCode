package single_number

import "testing"

func TestSingleNumber(t *testing.T) {
	if singleNumber([]int{2, 2, 1}) != 1 {
		t.Fatal("should be 1")
	}

	if singleNumber([]int{4, 1, 2, 1, 2}) != 4 {
		t.Fatal("should be 4")
	}
}
