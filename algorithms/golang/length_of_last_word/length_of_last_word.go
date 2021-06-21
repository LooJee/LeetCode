package length_of_last_word

func lengthOfLastWord(s string) int {
	length := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			length++
		} else if s[i] == ' ' && length != 0 {
			break
		}
	}

	return length
}
