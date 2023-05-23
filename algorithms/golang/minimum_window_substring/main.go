package minimumwindowsubstring

import (
	"math"
)

func minWindow(s string, t string) string {
	var (
		need   = make(map[byte]int)
		window = make(map[byte]int)
		found  = 0
		l, r   = 0, 0
		start  = 0
		length = math.MaxInt
		bs     = []byte(s)
		bt     = []byte(t)
	)

	for _, ele := range bt {
		need[ele]++
	}

	for r < len(bs) {
		c := bs[r]
		// 右边界右移
		r++
		// 判断是不是需要判断字串的元素
		if _, ok := need[c]; !ok {
			continue
		}

		window[c]++

		if window[c] == need[c] {
			found++
		}

		for found == len(need) {
			if length > r-l {
				start = l
				length = r - l
			}
			lc := bs[l]
			// 左边界收缩
			l++

			if _, ok := need[lc]; !ok {
				continue
			}

			if window[lc] == need[lc] {
				found--
			}

			window[lc]--
		}
	}

	if length == math.MaxInt {
		return ""
	} else {
		return string(bs[start : start+length])
	}
}
