package wordbreak

func wordBreak(s string, wordDict []string) bool {
	var (
		dict = make(map[string]bool)
		dp   = make([]bool, len(s))
	)

	for _, word := range wordDict {
		dict[word] = true
	}

	for i := 0; i < len(s); i++ {
		for j := i; j >= 0; j-- {
			if (j == 0 || dp[j-1]) && dict[s[j:i+1]] {
				dp[i] = true
			}
		}
	}

	return dp[len(s)-1]
}
