package lcof_reverse_left_words

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}
