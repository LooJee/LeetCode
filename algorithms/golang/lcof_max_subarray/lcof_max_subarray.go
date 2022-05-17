package lcof_max_subarray

import "math"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := math.MinInt
	curSum := maxSum
	for i := 0; i < len(nums); i++ {
		//如果之前的最大值小于0，和当前值相加的结果会比当前值更小，那么最大值可以直接设置为数组当前索引的值
		if curSum < 0 {
			curSum = nums[i]
		} else {
			curSum += nums[i]
		}
		maxSum = maxInt(maxSum, curSum)
	}

	return maxSum
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
