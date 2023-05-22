package searchrange

func searchRange(nums []int, target int) []int {
	l := leftBound(nums, target)
	r := rightBound(nums, target)

	return []int{l, r}
}

func leftBound(nums []int, target int) int {
	var (
		l, r = 0, len(nums) - 1
	)

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] == target {
			// r 最终位置会到 target 左侧边界的前一个位置
			r = mid - 1
		}
	}

	// 说明没有找到过值为 target 的元素
	if l == len(nums) {
		return -1
	}

	if nums[l] != target {
		return -1
	}

	return l
}

func rightBound(nums []int, target int) int {
	var (
		l, r = 0, len(nums) - 1
	)

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] == target {
			l = mid + 1
		}
	}

	if l == 0 {
		return -1
	}

	if nums[l-1] != target {
		return -1
	}

	return l - 1
}
