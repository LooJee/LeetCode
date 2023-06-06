package deleteandearn

import (
	"math"
)

func deleteAndEarn(nums []int) int {
	var (
		maxVal = math.MinInt
		arr    []int
	)
	// 构建可以动态规划的数组
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}

	arr = make([]int, maxVal+1)

	for _, num := range nums {
		arr[num]++
	}

	//开始动态规划
	var (
		a, b = arr[1], arr[0]
	)

	for i := 2; i < maxVal+1; i++ {
		a, b = int(math.Max(float64(a), float64(b+arr[i]*i))), a
	}

	return a
}
