package besttimetobuyandsellstockiv

import "math"

func maxProfit(k int, prices []int) int {
	var dp = make([][][]int, len(prices))

	for i := 0; i < len(prices); i++ {
		dp[i] = make([][]int, k+1)
		for theK := 0; theK <= k; theK++ {
			dp[i][theK] = make([]int, 2)
		}
	}

	for i := 0; i < len(prices); i++ {
		for kk := 1; kk <= k; kk++ {
			if i-1 < 0 {
				// 初始化
				dp[i][kk][0] = 0
				dp[i][kk][1] = -prices[i]
				continue
			}
			dp[i][kk][0] = int(math.Max(float64(dp[i-1][kk][0]), float64(dp[i-1][kk][1]+prices[i])))
			dp[i][kk][1] = int(math.Max(float64(dp[i-1][kk][1]), float64(dp[i-1][kk-1][0]-prices[i])))
		}
	}

	return dp[len(prices)-1][k][0]
}
