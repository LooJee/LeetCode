package lcof_reverse_words

import "strings"

func reverseWords(s string) string {
	r := len(s) - 1

	b := strings.Builder{}
	i := 0
	first := true
	for r >= 0 {
		if s[r] == ' ' && i > 0 {
			if !first {
				b.WriteByte(' ')
			}
			b.WriteString(s[r+1 : r+i+1])
			i = 0
			first = false
		} else if s[r] != ' ' {
			i++
		}
		r--
	}

	if i > 0 {
		if !first {
			b.WriteByte(' ')
		}
		b.WriteString(s[r+1 : r+i+1])
	}

	return b.String()
}
