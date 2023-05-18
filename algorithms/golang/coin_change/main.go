package coinchange

import "math"

func coinChange(coins []int, amount int) int {
	memos := make([]int, amount+1)
	for idx := range memos {
		memos[idx] = -2
	}
	return dp(coins, memos, amount)
}

func dp(coins, memos []int, amount int) int {
	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}

	if memos[amount] != -2 {
		return memos[amount]
	}

	minCoins := math.MaxInt

	for _, coin := range coins {
		sub := amount - coin
		ret := dp(coins, memos, sub)
		if ret == -1 {
			continue
		}

		minCoins = int(math.Min(float64(minCoins), float64(ret+1)))
	}

	if minCoins == math.MaxInt {
		memos[amount] = -1
		return -1
	} else {
		memos[amount] = minCoins
	}

	return memos[amount]
}
