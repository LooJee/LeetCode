package PascalsTriangleII

/*
如果不考虑空间复杂度，那么可以直接使用PascalsTriangle的方法来计算第k行的值，下面给出空间复杂度为O(k)的解法
*/
func getRow(rowIndex int) []int {
	b := make([]int, rowIndex+1)


	y := true

	for i := 0; i <= rowIndex ; i++ {
		if y {
			b[0] = 1
			for j := 1; j < rowIndex/2 + 1; j++ {
				b[j] = b[rowIndex-j+1] + b[rowIndex-j]
			}
		} else {
			b[rowIndex] = 1
			for j := 1; j < rowIndex/2 + 1; j++ {
				b[rowIndex-j] = b[j-1] + b[j]
			}
		}
		y = !y
	}

	if !y {
		for i := 0; i <= rowIndex/2; i++ {
			b[rowIndex-i] = b[i]
		}
	} else {
		for i := 0; i <= rowIndex/2; i++ {
			b[i] = b[rowIndex-i]
		}
	}

	return b
}
