package lcof_first_uniq_char

func firstUniqChar(s string) byte {
	chars := [26]int{}

	for i := 0; i < len(s); i++ {
		chars[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if chars[s[i]-'a'] == 1 {
			return s[i]
		}
	}

	return ' '
}
