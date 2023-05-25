package besttimetobuyandsellstock

import "math"

func maxProfit(prices []int) int {
	var (
		dp_i_0, dp_i_1 = 0, math.MinInt //dp_i_0 表示今日不持有股票，dp_i_1 表示今日持有股票
	)

	for i := 0; i < len(prices); i++ {
		dp_i_0 = int(math.Max(float64(dp_i_0), float64(dp_i_1+prices[i])))
		dp_i_1 = int(math.Max(float64(dp_i_1), float64(-prices[i])))
	}

	// 返回不持有时的总金额
	return dp_i_0
}
