package miceandcheese

import "sort"

// 要知道第一只老鼠恰好吃掉 k 口的最大值，可以假设都由老鼠2吃掉所有的奶酪，
// 然后计算 []diff =  []reward1 - []reward2， 得到奖励差，然后将 diff 排序，倒序取出差值最大的k个元素加到结果中
func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	var (
		reward int
		diff   = make([]int, len(reward1))
	)

	for i := 0; i < len(reward1); i++ {
		diff[i] = reward1[i] - reward2[i]
		reward += reward2[i]
	}

	sort.Ints(diff)

	for i := len(diff) - 1; i >= len(diff)-k; i-- {
		reward += diff[i]
	}

	return reward
}
