package summary_ranges

import "fmt"

/*
给定一个  无重复元素 的 有序 整数数组 nums 。

返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表 。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，
并且不存在属于某个范围但不属于 nums 的数字 x 。

列表中的每个区间范围 [a,b] 应该按如下格式输出：

"a->b" ，如果 a != b
"a" ，如果 a == b
*/
func summaryRanges(nums []int) []string {
	result := make([]string, 0)

	for i := 0; i < len(nums); {
		j := i + 1
		for ; j < len(nums) && nums[j]-1 == nums[j-1]; j++ {
		}

		if j-1 == i {
			result = append(result, fmt.Sprintf("%d", nums[i]))
		} else {
			result = append(result, fmt.Sprintf("%d->%d", nums[i], nums[j-1]))
		}
		i = j
	}

	return result
}
