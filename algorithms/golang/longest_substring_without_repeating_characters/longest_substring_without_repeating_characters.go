package longestsubstringwithoutrepeatingcharacters

import "math"

func lengthOfLongestSubstring(s string) int {
	var size int
	start := 0
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if preIdx, ok := m[s[i]]; ok {
			if start < preIdx {
				start = preIdx
			}
		}

		size = int(math.Max(float64(size), float64(i-start+1)))

		m[s[i]] = i + 1
	}

	return size
}
