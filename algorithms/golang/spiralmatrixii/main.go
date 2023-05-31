package spiralmatrixii

func generateMatrix(n int) [][]int {
	var (
		result = make([][]int, n)
		top    = 0
		bottom = n - 1
		left   = 0
		right  = n - 1
		value  = 1
	)

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	for top <= bottom && left <= right {
		if top <= bottom {
			for i := left; i <= right; i++ {
				result[top][i] = value
				value++
			}
			top++
		}

		if left <= right {
			for i := top; i <= bottom; i++ {
				result[i][right] = value
				value++
			}
			right--
		}

		if top <= bottom {
			for i := right; i >= left; i-- {
				result[bottom][i] = value
				value++
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				result[i][left] = value
				value++
			}
			left++
		}
	}

	return result
}
