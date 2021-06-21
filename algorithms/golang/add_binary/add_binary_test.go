package add_binary

import "testing"

func TestAddBinary(t *testing.T) {
	if addBinary("0", "1") != "1" {
		t.Fatal("should be 1")
	}

	if addBinary("1", "1") != "10" {
		t.Fatal("should be 10")
	}

	if addBinary("1", "10") != "11" {
		t.Fatal("should be 11")
	}

	if addBinary("1010", "1011") != "10101" {
		t.Fatal("should be 10101")
	}

	if addBinary("1", "111") != "1000" {
		t.Fatal("should be 1000")
	}

	if addBinary("101111", "10") != "110001" {
		t.Fatal("should be 110001")
	}
}
