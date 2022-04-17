package longestsubstringwithoutrepeatingcharacters

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	if lengthOfLongestSubstring("abcabcbb") != 3 {
		t.Fatal("should be 3")
	}

	if lengthOfLongestSubstring("bbbbb") != 1 {
		t.Fatal("should be 1")
	}

	if lengthOfLongestSubstring("pwwkew") != 3 {
		t.Fatal("should be 3")
	}
}
