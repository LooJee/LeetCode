package subsetsii

import "sort"

// 因为 nums 中有重复元素，回溯时需要剪枝，避免相同元素出现，例如 [1,2,2] ，剪枝后，第二个 2 的分支会被过滤掉
// 要实现剪枝，比较方便的方法是将 nums 排序之后，对比 nums[i] 和 nums[i-1] 是否相同，如果相同就跳过。
func subsetsWithDup(nums []int) [][]int {
	var (
		result = make([][]int, 0)
		track  = make([]int, 0)
	)

	sort.Ints(nums)

	backtrack(nums, track, &result, 0)

	return result
}

func backtrack(nums, track []int, result *[][]int, startIdx int) {
	tmp := make([]int, len(track))
	copy(tmp, track)
	*result = append(*result, tmp)

	for i := startIdx; i < len(nums); i++ {
		// 过滤掉同一级中的相同边
		if i > startIdx && nums[i] == nums[i-1] {
			continue
		}
		track = append(track, nums[i])
		backtrack(nums, track, result, i+1)
		track = track[:len(track)-1]
	}
}
