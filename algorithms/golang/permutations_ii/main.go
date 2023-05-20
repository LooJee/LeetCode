package permutationsii

import "sort"

// 采用回溯法解决问题，因为需要返回不重复的全排列，所以需要做剪枝
func permuteUnique(nums []int) [][]int {
	var (
		result = make([][]int, 0)
		track  = make([]int, 0)
		used   = make([]bool, len(nums))
	)

	sort.Ints(nums)
	backtrack(&result, nums, track, used)

	return result
}

func backtrack(result *[][]int, nums, track []int, used []bool) {
	if len(track) == len(nums) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		*result = append(*result, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		// 剪枝
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}

		track = append(track, nums[i])
		used[i] = true
		backtrack(result, nums, track, used)
		track = track[:len(track)-1]
		used[i] = false
	}
}
