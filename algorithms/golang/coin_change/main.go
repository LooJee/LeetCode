package coinchange

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}

	// base case
	dp[0] = 0

	// 自底向上求值
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}

			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coin]+1)))
		}
	}

	if dp[amount] == amount+1 {
		dp[amount] = -1
	}

	return dp[amount]
}
