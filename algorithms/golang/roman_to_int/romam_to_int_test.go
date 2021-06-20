package roman_to_int

import "testing"

func TestRomanToInt(t *testing.T) {
	if romanToInt("III") != 3 {
		t.Fatal("III should be 3")
	}

	if romanToInt("I") != 1 {
		t.Fatal("I should be 1")
	}

	if romanToInt("IV") != 4 {
		t.Fatal("IV should be 4")
	}

	if romanToInt("LVIII") != 58 {
		t.Fatal("LVIII should be 58")
	}
}
