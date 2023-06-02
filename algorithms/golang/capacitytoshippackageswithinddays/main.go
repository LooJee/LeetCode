package capacitytoshippackageswithinddays

func shipWithinDays(weights []int, days int) int {
	var (
		left, right = 0, 0
	)

	// 将问题演化为二分查找，左边界为 weights 中的最大重量，即一次至少能装载一个货物，
	// 右边界为所有货物的总重量，即一次将所有货物运走
	for _, weight := range weights {
		if weight > left {
			left = weight
		}

		right += weight
	}

	// 通过二分查找，找到最小能实现在 days 天内完成货物运送的承载量
	for left <= right {
		var (
			mid = left + (right-left)/2
			day = 0
			sum = 0
		)

		for i := 0; i < len(weights); i++ {
			sum += weights[i]
			if sum >= mid {
				day++
				if sum > mid {
					i--
				}
				sum = 0
			}
		}

		if sum != 0 {
			day++
		}

		if day > days {
			// 使用的天数比要求的多，说明设定的载货量太少了，需要增加
			left = mid + 1
		} else if day < days {
			// 如果使用的天数比要求的少，说明设定的载货量太多了，需要减少
			right = mid - 1
		} else if day == days {
			// 如果使用的天数和要求的相同，则减少载货量，找到最小载货量
			right = mid - 1
		}
	}

	return left
}
