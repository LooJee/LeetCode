package kokoeatingbananas

func minEatingSpeed(piles []int, h int) int {
	var (
		left, right = 1, 1
	)

	// 一次最多只能吃一堆，一堆的香蕉可以分多次吃，即最大值是数量最多的一堆
	for _, pile := range piles {
		if pile > right {
			right = pile
		}
	}

	for left <= right {
		var (
			mid   = left + (right-left)/2
			times = 0 // 次数
		)

		for i := 0; i < len(piles); i++ {
			times += piles[i] / mid

			if piles[i]%mid > 0 {
				times += 1
			}
		}

		if times > h {
			// 如果次数比要求的大，表示每次吃的太少了，需要增加，将左边界右移
			left = mid + 1
		} else if times < h {
			// 如果次数比要求的小，表示每次吃的有点多，需要减少，将右边界左移
			right = mid - 1
		} else if times == h {
			// 如果次数和要求的相同，则继续将右边界左移，找每次吃的最少量
			right = mid - 1
		}
	}

	return left
}
