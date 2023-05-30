package carpooling

import "fmt"

func carPoolingv2(trips [][]int, capacity int) bool {
	var (
		nums = make([]int, 1000)
	)

	// 构造差分数组
	for _, trip := range trips {
		nums[trip[1]] += trip[0]
		if trip[2] < len(nums) {
			nums[trip[2]] -= trip[0]
		}
	}

	fmt.Println(nums)

	// 还原
	if nums[0] > capacity {
		return false
	}
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
		if nums[i] > capacity {
			return false
		}
	}

	return true
}
