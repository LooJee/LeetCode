package permutationinstring

func checkInclusion(s1 string, s2 string) bool {
	var (
		need   = make(map[byte]int)
		window = make(map[byte]int)
		l, r   = 0, 0
		b1, b2 = []byte(s1), []byte(s2)
		found  = 0
	)

	for _, ele := range b1 {
		need[ele]++
	}

	for r < len(b2) {
		c := b2[r]
		r++

		if _, ok := need[c]; ok {
			window[c]++

			if window[c] == need[c] {
				found++
			}
		}

		// 将窗口大小保持为 b1 的大小，判断窗口内的元素是否是 b1 的排列
		if r-l >= len(b1) {
			if found == len(need) {
				return true
			}

			lc := b2[l]
			l++

			if _, ok := need[lc]; ok {
				if window[lc] == need[lc] {
					found--
				}

				window[lc]--
			}
		}
	}

	return false
}
