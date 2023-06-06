package equalrowandcolumnpairs

func equalPairs(grid [][]int) int {
	var (
		cols  = make([]int, len(grid[0]))
		rows  = make([]int, len(grid))
		cnt   = 0
		memos = make(map[int]int)
	)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			rows[i] = rows[i]*10 + grid[i][j]
			cols[j] = cols[j]*10 + grid[i][j]
		}
	}

	for i := 0; i < len(rows); i++ {
		memos[rows[i]]++
	}

	for i := 0; i < len(cols); i++ {
		cnt += memos[cols[i]]
	}

	return cnt		
}
