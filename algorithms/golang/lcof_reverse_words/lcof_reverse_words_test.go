package lcof_reverse_words

import "testing"

func TestReverseWords(t *testing.T) {
	t.Log(reverseWords("the sky is blue"))
	t.Log(reverseWords("  hello world!  "))
}
