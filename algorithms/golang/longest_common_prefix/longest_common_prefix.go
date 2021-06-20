package longest_common_prefix

import "math"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minLen := math.MaxInt32
	for i := range strs {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
		}
	}

	prefix := strs[0][:minLen]
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			if prefix[j] != strs[i][j] {
				prefix = prefix[:j]
				continue
			}
		}
	}

	return prefix
}
