package lcof_max_value

func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	dp := grid

	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ {
			if j == 0 && i == 0 {
				continue
			} else if j == 0 && i != 0 {
				dp[j][i] = grid[j][i] + dp[j][i-1]
			} else if j != 0 && i == 0 {
				dp[j][i] = grid[j][i] + dp[j-1][i]
			} else {
				dp[j][i] = grid[j][i] + maxInt(dp[j-1][i], dp[j][i-1])
			}
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}
