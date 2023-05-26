package houserobber

import "math"

func rob(nums []int) int {
	var (
		dp_i_1, dp_i_2 = 0, 0 // dp_i_1 = dp[i-1], dp_i_2 = dp[i-2]
	)

	for i := 0; i < len(nums); i++ {
		dp_i_1, dp_i_2 = int(math.Max(float64(dp_i_1), float64(dp_i_2+nums[i]))), dp_i_1
	}

	return dp_i_1
}
