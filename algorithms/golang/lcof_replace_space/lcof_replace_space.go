package lcof_replace_space

import "strings"

func replaceSpace(s string) string {
	const theSpace = "%20"
	builder := strings.Builder{}

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			builder.WriteByte(s[i])
		} else {
			builder.WriteString(theSpace)
		}
	}

	return builder.String()
}
