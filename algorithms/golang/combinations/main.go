package combinations

func combine(n int, k int) [][]int {
	var (
		result = make([][]int, 0)
		track  = make([]int, 0)
	)

	backtrack(n, k, 1, &result, track)

	return result
}

func backtrack(n, k, startNum int, result *[][]int, track []int) {
	if len(track) == k {
		tmp := make([]int, k)
		copy(tmp, track)
		*result = append(*result, tmp)

		return
	}

	for i := startNum; i <= n; i++ {
		track = append(track, i)
		backtrack(n, k, i+1, result, track)
		track = track[:len(track)-1]
	}
}
