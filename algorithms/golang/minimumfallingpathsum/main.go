package minimumfallingpathsum

import "math"

func minFallingPathSum(matrix [][]int) int {
	var (
		rowSize = len(matrix)
		colSize = len(matrix[0])
		dp      = make([][]int, len(matrix))
	)

	for i := 0; i < rowSize; i++ {
		dp[i] = make([]int, colSize)
	}

	for i := 0; i < colSize; i++ {
		dp[0][i] = matrix[0][i]
	}

	for i := 1; i < rowSize; i++ {
		for j := 0; j < colSize; j++ {
			minVal := dp[i-1][j]
			if j-1 >= 0 {
				minVal = int(math.Min(float64(minVal), float64(dp[i-1][j-1])))
			}

			if j+1 < colSize {
				minVal = int(math.Min(float64(minVal), float64(dp[i-1][j+1])))
			}

			dp[i][j] = minVal + matrix[i][j]
		}
	}

	return findMinVal(dp[rowSize-1])
}

func findMinVal(nums []int) int {
	minVal := math.MaxInt

	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
	}

	return minVal
}
