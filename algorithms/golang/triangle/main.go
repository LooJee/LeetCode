package triangle

import "math"

func minimumTotal(triangle [][]int) int {
	var (
		rowSize = len(triangle)
		dp      = make([][]int, rowSize)
	)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(triangle[i]))
	}

	for i := 0; i < len(dp[rowSize-1]); i++ {
		dp[rowSize-1][i] = triangle[rowSize-1][i]
	}

	for i := rowSize - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[i][j] = triangle[i][j] + int(math.Min(float64(dp[i+1][j]), float64(dp[i+1][j+1])))
		}
	}

	return dp[0][0]
}
