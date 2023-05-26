package houserobberii

import "math"

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return int(math.Max(float64(doRob(nums[:len(nums)-1])), float64(doRob(nums[1:]))))
}

func doRob(nums []int) int {
	var (
		dp_i_1, dp_i_2 = 0, 0
	)

	for i := 0; i < len(nums); i++ {
		dp_i_1, dp_i_2 = int(math.Max(float64(dp_i_1), float64(dp_i_2+nums[i]))), dp_i_1
	}

	return dp_i_1
}
