package PascalsTriangle

func generate(numRows int) [][]int {
	r := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		a := make([]int, i+1)
		a[0] = 1
		a[i] = 1

		for j := 1; j < i; j++ {
			a[j] = r[i-1][j-1]+r[i-1][j]
		}

		r[i] = a
	}

	return r
}
