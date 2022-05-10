package valid_perfect_square

import "testing"

func TestIsPerfectSquare(t *testing.T) {
	if !isPerfectSquare(16) {
		t.Fatal("should be true")
	}

	if !isPerfectSquare(1) {
		t.Fatal("should be true")
	}
}
