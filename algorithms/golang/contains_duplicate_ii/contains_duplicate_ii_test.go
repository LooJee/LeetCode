package contains_duplicate_ii

import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	if !containsNearbyDuplicate([]int{1, 2, 3, 1}, 3) {
		t.Fatal("should be true")
	}
}
