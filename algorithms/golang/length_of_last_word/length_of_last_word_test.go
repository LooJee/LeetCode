package length_of_last_word

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	if lengthOfLastWord(" ") != 0 {
		t.Fatal("should be 0")
	}

	if lengthOfLastWord("hello") != 5 {
		t.Fatal("should be 5")
	}

	if lengthOfLastWord("hello world ") != 5 {
		t.Fatal("should be 5")
	}

	if lengthOfLastWord("hello world") != 5 {
		t.Fatal("should be 5")
	}
}
