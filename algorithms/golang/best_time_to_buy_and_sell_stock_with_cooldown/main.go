package besttimetobuyandsellstockwithcooldown

import "math"

func maxProfit(prices []int) int {
	var (
		dp_i_0, dp_i_1 = 0, math.MinInt
		dp_cooldown    = 0
	)

	for i := 0; i < len(prices); i++ {
		tmp := dp_i_0 // 缓存昨天未持有股票时的持有金额
		dp_i_0 = int(math.Max(float64(dp_i_0), float64(dp_i_1+prices[i])))
		dp_i_1 = int(math.Max(float64(dp_i_1), float64(dp_cooldown-prices[i])))
		// 这个循环中缓存昨天未持有股票的持有金额，下一个循环中就会成为前天未持有股票时的持有金额
		dp_cooldown = tmp
	}

	return dp_i_0
}
