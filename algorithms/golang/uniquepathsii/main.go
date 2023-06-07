package uniquepathsii

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var (
		rowSize = len(obstacleGrid)
		colSize = len(obstacleGrid[0])
		dp      = make([][]int, rowSize+1)
	)

	for i := 0; i <= rowSize; i++ {
		dp[i] = make([]int, colSize+1)
	}

	for i := 1; i <= rowSize; i++ {
		for j := 1; j <= colSize; j++ {
			if i == 1 && j == 1 && obstacleGrid[i-1][i-1] == 0 {
				dp[i][j] = 1
				continue
			}

			if obstacleGrid[i-1][j-1] == 1 {
				dp[i][j] = 0
				continue
			}

			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[rowSize][colSize]
}
