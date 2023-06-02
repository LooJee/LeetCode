package splitarraylargestsum

func splitArray(nums []int, k int) int {
	var (
		left, right = 0, 0
	)

	for _, num := range nums {
		if num > left {
			left = num
		}

		right += num
	}

	for left <= right {
		var (
			mid = left + (right-left)/2
			sum = 0
			cnt = 0
		)

		for i := 0; i < len(nums); i++ {
			sum += nums[i]

			if sum >= mid {
				if sum > mid {
					i--
				}
				cnt++
				sum = 0
			}
		}
		if sum != 0 {
			cnt++
		}

		if cnt > k {
			// 子数组数量比要求的多，增大 mid， 左边界右移
			left = mid + 1
		} else if cnt <= k {
			// 子数组数量比要求的少，减小 mid，右边界左移
			right = mid - 1
		} else if cnt == k {
			// 子数组数量和要求的相同，减小 mid，求最小最大值
			right = mid - 1
		}
	}

	return left
}
