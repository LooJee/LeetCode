package lcof_length_oflongest_substring

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Log(lengthOfLongestSubstring("abcabcbb"))
	t.Log(lengthOfLongestSubstring("bbbbbbb"))
	t.Log(lengthOfLongestSubstring("a"))
	t.Log(lengthOfLongestSubstring("aab"))
	t.Log(lengthOfLongestSubstring("abba"))
	t.Log(lengthOfLongestSubstring("avab"))
}
