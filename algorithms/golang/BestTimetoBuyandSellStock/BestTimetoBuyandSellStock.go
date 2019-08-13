package BestTimetoBuyandSellStock
/*
	时间复杂度为 O(n²)，简单的在循环中找到最大和
*/
func maxProfit1(prices []int) int {
	max := 0

	for i := 0; i < len(prices); i++ {
		for j := i+1; j < len(prices); j++ {
			k := prices[j] - prices[i]
			if k > max {
				max = k
			}
		}
	}

	return max
}

/*
	时间复杂度为O(n)，在循环中找到数组中的最小值，并将之后的值与它计算，如果结果大于保存的最大值，则将该结果置为最大值
*/
func maxProfit2(prices []int) int {
	max := 0
	var min int

	if len(prices) != 0 {
		min = prices[0]
	} else {
		return max
	}

	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			if prices[i] - min > max {
				max = prices[i] - min
			}
		}
	}

	return max
}
