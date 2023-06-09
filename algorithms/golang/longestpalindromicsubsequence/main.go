package longestpalindromicsubsequence

import "math"

func longestPalindromeSubseq(s string) int {
	var (
		dp = make([][]int, len(s))
	)

	for i := range dp {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j <= len(s)-1; j++ {
			if s[i] != s[j] {
				dp[i][j] = int(math.Max(float64(dp[i+1][j]), float64(dp[i][j-1])))
			} else {
				dp[i][j] = dp[i+1][j-1] + 2
			}
		}
	}

	return dp[0][len(s)-1]
}
