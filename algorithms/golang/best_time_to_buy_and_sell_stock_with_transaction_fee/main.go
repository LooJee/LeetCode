package besttimetobuyandsellstockwithtransactionfee

import "math"

func maxProfit(prices []int, fee int) int {
	var (
		dp_i_0, dp_i_1 = 0, math.MinInt
	)

	for i := 0; i < len(prices); i++ {
		tmp := dp_i_0
		dp_i_0 = int(math.Max(float64(dp_i_0), float64(dp_i_1+prices[i])))
		dp_i_1 = int(math.Max(float64(dp_i_1), float64(tmp-prices[i]-fee)))
	}

	return dp_i_0
}
