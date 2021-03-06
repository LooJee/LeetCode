package BestTimetoBuyandSellStockII

func maxProfit(prices []int) int {
	max := 0

	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			max += prices[i] - prices[i-1]
		}
	}

	return max
}
