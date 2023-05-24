package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	var (
		l, r   = 0, 0
		length = 0
		window = make(map[byte]int)
	)

	for r < len(s) {
		rc := s[r]
		r++
		window[rc]++

		for window[rc] > 1 {
			window[s[l]]--
			l++
		}

		if r-l > length {
			length = r - l
		}
	}

	return length
}
