package lcof_max_profit

import "math"

func maxProfit(prices []int) int {
	maxRes := 0

	min := math.MaxInt
	for i := 0; i < len(prices); i++ {
		curProfit := prices[i] - min
		if maxRes < curProfit {
			maxRes = curProfit
		}

		if min > prices[i] {
			min = prices[i]
		}
	}

	return maxRes
}
