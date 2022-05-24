package lcof_length_oflongest_substring

import "math"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	m := make(map[byte]int)
	longest := math.MinInt

	l := 0
	for i := 0; i < len(s); i++ {
		if c, ok := m[s[i]]; ok && c >= l {
			tmp := i - l
			if tmp > longest {
				longest = tmp
			}
			l = c + 1
		}
		m[s[i]] = i
	}

	if len(s)-l > longest {
		longest = len(s) - l
	}

	return longest
}
