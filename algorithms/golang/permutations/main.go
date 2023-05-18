package permutations

func permute(nums []int) [][]int {
	var (
		used   = make([]bool, len(nums))
		result = make([][]int, 0)
		track  = make([]int, 0)
	)

	backtrack(nums, used, track, &result)

	return result
}

func backtrack(nums []int, used []bool, track []int, result *[][]int) {
	if len(track) == len(nums) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		*result = append(*result, tmp)
	}

	for idx, num := range nums {
		if used[idx] {
			continue
		}
		used[idx] = true
		track = append(track, num)
		backtrack(nums, used, track, result)
		used[idx] = false
		track = track[:len(track)-1]
	}
}
