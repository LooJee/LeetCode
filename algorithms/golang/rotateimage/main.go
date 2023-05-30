package rotateimage

// 两次翻转，先沿斜对角一次翻转，再沿中间一次翻转
func rotate(matrix [][]int) {
	diagonalRotate(matrix)
	verticalRotate(matrix)
}

func diagonalRotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix[0]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func verticalRotate(matrix [][]int) {
	middle := len(matrix[0]) >> 1
	last := len(matrix[0]) - 1

	if len(matrix[0])%2 != 0 {
		middle++
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < middle; j++ {
			matrix[i][j], matrix[i][last-j] = matrix[i][last-j], matrix[i][j]
		}
	}
}
