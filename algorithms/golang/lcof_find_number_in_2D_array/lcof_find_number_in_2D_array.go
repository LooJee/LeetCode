package lcof_find_number_in_2D_array

//因为二维数组是有序的，所以可以从右上角开始进行搜索
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	i := len(matrix[0]) - 1
	j := 0

	for i >= 0 && j < len(matrix) {
		if matrix[j][i] == target {
			return true
		} else if matrix[j][i] < target {
			j++
		} else {
			i--
		}
	}

	return false
}
