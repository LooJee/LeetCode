package valid_perfect_square

func isPerfectSquare(num int) bool {
	l, e := 0, num
	var mid int
	for l <= e {
		mid = (l + e) / 2
		t := mid * mid
		if t < num {
			l = mid + 1
		} else if t > num {
			e = mid - 1
		} else {
			return true
		}
	}

	return false
}
