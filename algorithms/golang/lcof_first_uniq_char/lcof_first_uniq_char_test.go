package lcof_first_uniq_char

import "testing"

func TestFirstUniqChar(t *testing.T) {
	if firstUniqChar("") != ' ' {
		t.Fatal("should be ' '")
	}

	if firstUniqChar("abaccdeff") != 'b' {
		t.Fatal("should be 'b'")
	}
}
