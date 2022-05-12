package search_range

import "math"

func searchRange(nums []int, target int) []int {
	l := searchLeft(nums, target)
	r := searchRight(nums, target)

	return []int{l, r}
}

func searchLeft(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	pos := -1

	var mid int
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else {
			if nums[mid] == target {
				pos = mid
			}
			r = mid - 1
		}
	}

	return pos
}

func searchRight(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	pos := -1

	var mid int
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] > target {
			r = mid - 1
		} else {
			if nums[mid] == target {
				pos = mid
			}
			l = mid + 1
		}
	}

	return pos
}

func searchRangeOnce(nums []int, target int) []int {
	l, r := math.MaxInt, -1

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			if i < l {
				l = i
			}

			if i > r {
				r = i
			}
		}
	}

	if l == math.MaxInt {
		l = -1
	}

	return []int{l, r}
}
