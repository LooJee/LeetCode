package combinationsum

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	var (
		result = make([][]int, 0)
		track  = make([]int, 0)
	)

	sort.Ints(candidates)
	backtrack(candidates, track, &result, target, 0)

	return result
}

func backtrack(candidates, track []int, result *[][]int, target, startIdx int) {
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
		if i != startIdx && candidates[i] == candidates[i-1] {
			continue
		}

		track = append(track, candidates[i])
		// startIdx 传递 i ，表示同一个分支下元素可以重复
		backtrack(candidates, track, result, target-candidates[i], i)
		track = track[:len(track)-1]
	}
}
