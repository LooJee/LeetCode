package minimumpathsum

import (
	"math"
)

func minPathSum(grid [][]int) int {
	var (
		rowSize = len(grid)
		colSize = len(grid[0])
		dp      = make([][]int, rowSize+1)
	)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, colSize+1)
		dp[i][0] = math.MaxInt
	}

	for i := 0; i <= colSize; i++ {
		dp[0][i] = math.MaxInt
	}

	dp[0][1], dp[1][0] = 0, 0

	for i := 1; i <= rowSize; i++ {
		for j := 1; j <= colSize; j++ {
			dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1]))) + grid[i-1][j-1]
		}
	}

	return dp[rowSize][colSize]
}
