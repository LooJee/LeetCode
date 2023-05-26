package besttimetobuyandsellstockiii

import "math"

func maxProfit(prices []int) int {
	var dp [][][]int
	const K = 2

	dp = make([][][]int, len(prices))

	for i := 0; i < len(prices); i++ {
		dp[i] = make([][]int, K+1)
		for k := 0; k <= K; k++ {
			dp[i][k] = make([]int, K)
		}
	}

	for i := 0; i < len(prices); i++ {
		for k := 1; k <= K; k++ {
			if i-1 < 0 {
				//初始化
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}

			dp[i][k][0] = int(math.Max(float64(dp[i-1][k][0]), float64(dp[i-1][k][1]+prices[i])))
			dp[i][k][1] = int(math.Max(float64(dp[i-1][k][1]), float64(dp[i-1][k-1][0]-prices[i])))
		}
	}

	return dp[len(prices)-1][K][0]
}

func maxProfit2(prices []int) int {
	var (
		dp_i_k1_0, dp_i_k1_1, dp_i_k2_0, dp_i_k2_1 = 0, math.MinInt, 0, math.MinInt
	)

	for i := 0; i < len(prices); i++ {
		dp_i_k2_0 = int(math.Max(float64(dp_i_k2_0), float64(dp_i_k2_1+prices[i])))
		dp_i_k2_1 = int(math.Max(float64(dp_i_k2_1), float64(dp_i_k1_0-prices[i])))
		dp_i_k1_0 = int(math.Max(float64(dp_i_k1_0), float64(dp_i_k1_1+prices[i])))
		dp_i_k1_1 = int(math.Max(float64(dp_i_k1_1), float64(-prices[i])))
	}

	return dp_i_k2_0
}
