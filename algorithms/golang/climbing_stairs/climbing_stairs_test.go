package climbing_stairs

import "testing"

func TestClimbStairs(t *testing.T) {
	if climbStairs(1) != 1 {
		t.Fatal("should be 1")
	}

	if climbStairs(2) != 2 {
		t.Fatal("should be 2")
	}

	if climbStairs(3) != 3 {
		t.Fatal("should be 3")
	}

	if climbStairs(4) != 5 {
		t.Fatal("should be 5")
	}
}
