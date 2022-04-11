package majority_element

import "testing"

func TestMajorityElement(t *testing.T) {
	if majorityElement([]int{3, 2, 3}) != 3 {
		t.Fatal("should be 3")
	}
}
