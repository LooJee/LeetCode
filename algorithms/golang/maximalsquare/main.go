package maximalsquare

import (
	"math"
)

func maximalSquare(matrix [][]byte) int {
	var (
		rowSize = len(matrix)
		colSize = len(matrix[0])
		dp      = make([][]int, rowSize) // dp 记录边长
		maxEdge = math.MinInt
	)

	for i := 0; i < rowSize; i++ {
		dp[i] = make([]int, colSize)
	}

	for i := 0; i < rowSize; i++ {
		for j := 0; j < colSize; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			if i == 0 || j == 0 {
				dp[i][j] = int(matrix[i][j] - '0')
			} else {
				dp[i][j] = int(math.Min(math.Min(float64(dp[i-1][j-1]), float64(dp[i-1][j])), float64(dp[i][j-1]))) + 1
			}
			if dp[i][j] > maxEdge {
				maxEdge = dp[i][j]
			}
		}
	}

	return maxEdge * maxEdge
}
