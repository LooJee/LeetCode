package sqrtx

import "testing"

func TestMySqrt(t *testing.T) {
	if mySqrt(4) != 2 {
		t.Fatal("should be 2")
	}

	if mySqrt(5) != 2 {
		t.Fatal("should be 2")
	}

	if mySqrt(8) != 2 {
		t.Fatal("should be 2")
	}

	if mySqrt(9) != 3 {
		t.Fatal("should be 2")
	}
}
