package subsets

func subsets(nums []int) [][]int {
	var (
		result = make([][]int, 0)
		track  = make([]int, 0)
	)

	backtrack(0, &result, track, nums)

	return result
}

func backtrack(startIdx int, result *[][]int, track, nums []int) {
	tmp := make([]int, len(track))
	copy(tmp, track)
	*result = append(*result, tmp)

	for i := startIdx; i < len(nums); i++ {
		track = append(track, nums[i])
		backtrack(i+1, result, track, nums)
		track = track[:len(track)-1]
	}
}
