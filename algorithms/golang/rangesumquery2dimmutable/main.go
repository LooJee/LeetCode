package rangesumquery2dimmutable

type NumMatrix struct {
	preSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	var (
		row, column = len(matrix), len(matrix[0])
		numMatrix   = NumMatrix{preSum: make([][]int, row+1)}
	)

	for i := 0; i <= row; i++ {
		numMatrix.preSum[i] = make([]int, column)
	}

	for i := 1; i <= row; i++ {
		for j := 1; j <= column; j++ {
			numMatrix.preSum[i][j] = numMatrix.preSum[i-1][j] + numMatrix.preSum[i][j-1] + matrix[i-1][j-1] - numMatrix.preSum[i-1][j-1]
		}
	}

	return numMatrix
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSum[row2+1][col2+1] - this.preSum[row2+1][col1] - this.preSum[row1][col2+1] + this.preSum[row1][col1]
}
