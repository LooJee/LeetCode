package reverse

import (
	"math"
	"testing"
)

func TestReverse(t *testing.T) {
	if reverse(-123) != -321 {
		t.Fatal("reverse -123 failed")
	}

	if reverse(math.MaxInt32+1) != 0 {
		t.Fatal("reverse max int32 failed")
	}

	if reverse(math.MinInt32-1) != 0 {
		t.Fatal("reverse min int32 failed")
	}
}

func TestReverse2(t *testing.T) {
	if reverse2(-123) != -321 {
		t.Fatal("reverse -123 failed")
	}

	if reverse2(math.MaxInt32+1) != 0 {
		t.Fatal("reverse max int32 failed")
	}

	if reverse2(math.MinInt32-1) != 0 {
		t.Fatal("reverse min int32 failed")
	}
}
