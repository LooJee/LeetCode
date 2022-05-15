package reverse_words_in_a_string

import "strings"

func reverseWords(s string) string {
	builder := strings.Builder{}
	for j := len(s) - 1; j >= 0; j-- {
		if s[j] == ' ' {
			continue
		} else {
			i := j
			for ; i >= 0 && s[i] != ' '; i-- {
			}
			i++
			if builder.Len() != 0 {
				builder.WriteByte(' ')
			}
			builder.WriteString(s[i : j+1])
			j = i
		}
	}

	return builder.String()
}
