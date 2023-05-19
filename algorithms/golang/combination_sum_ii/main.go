package combinationsumii

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var (
		result = make([][]int, 0)
		track  = make([]int, 0)
		used   = make(map[int]bool)
	)

	sort.Ints(candidates)
	backtrack(candidates, track, used, target, 0, &result)

	return result
}

func backtrack(candidates, track []int, used map[int]bool, target, startIdx int, result *[][]int) {
	if target < 0 {
		return
	}

	if target == 0 {
		tmp := make([]int, len(track))
		copy(tmp, track)
		*result = append(*result, tmp)
		return
	}

	for i := startIdx; i < len(candidates); i++ {
		if i > startIdx && candidates[i] == candidates[i-1] {
			continue
		}

		target -= candidates[i]
		track = append(track, candidates[i])
		backtrack(candidates, track, used, target, i+1, result)
		target += candidates[i]
		track = track[:len(track)-1]
	}
}
