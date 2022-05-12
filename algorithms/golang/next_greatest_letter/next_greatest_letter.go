package next_greatest_letter

func nextGreatestLetter(letters []byte, target byte) byte {
	l := 0
	r := len(letters) - 1
	for l <= r {
		mid := (l + r) / 2
		if letters[mid] <= target {
			l = mid + 1
		} else if letters[mid] > target && mid-1 >= 0 && letters[mid-1] <= target {
			return letters[mid]
		} else {
			r = mid - 1
		}
	}

	return letters[0]
}
