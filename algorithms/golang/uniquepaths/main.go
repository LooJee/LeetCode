package uniquepaths

func uniquePaths(m int, n int) int {
	var (
		dp = make([][]int, m+1)
	)

	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	dp[1][1] = 1

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 && j == 1 {
				continue
			}

			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m][n]
}
