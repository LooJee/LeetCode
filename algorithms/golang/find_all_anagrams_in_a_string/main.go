package findallanagramsinastring

func findAnagrams(s string, p string) []int {
	var (
		need   = make(map[byte]int)
		window = make(map[byte]int)
		l, r   = 0, 0
		found  = 0
		bs, bp = []byte(s), []byte(p)
		result = make([]int, 0)
	)

	for _, ele := range bp {
		need[ele]++
	}

	for r < len(bs) {
		rc := bs[r]
		r++

		if _, ok := need[rc]; ok {
			window[rc]++

			if window[rc] == need[rc] {
				found++
			}
		}

		if r-l == len(bp) {
			if found == len(need) {
				result = append(result, l)
			}

			lc := bs[l]
			l++

			if _, ok := need[lc]; ok {
				if window[lc] == need[lc] {
					found--
				}
				window[lc]--
			}
		}
	}

	return result
}
