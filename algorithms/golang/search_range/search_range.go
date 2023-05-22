package searchrange

import "math"

func searchRange(nums []int, target int) []int {
	var (
		l, r = math.MaxInt, -1
	)

	for i := 0; i < len(nums); i++ {
		if nums[i] == target && i < l {
			l = i
		}

		if nums[i] == target && i > r {
			r = i
		}
	}

	if l == math.MaxInt {
		l, r = -1, -1
	}

	return []int{l, r}
}
