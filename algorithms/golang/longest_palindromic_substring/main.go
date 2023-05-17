package longestpalindromicsubstring

func longestPalindrome(s string) string {
	longgest := ""
	for i := 0; i < len(s); i++ {
		s1 := findPalindrome(s, i, i)
		s2 := findPalindrome(s, i, i+1)

		if len(s1) > len(longgest) {
			longgest = s1
		}

		if len(s2) > len(longgest) {
			longgest = s2
		}
	}

	return longgest
}

func findPalindrome(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return s[left+1 : right]
}
