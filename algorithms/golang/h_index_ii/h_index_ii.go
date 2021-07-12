package h_index_ii

func hIndex(citations []int) int {
	left := 0
	right := len(citations) - 1

	for left <= right {
		mid := left + (right-left)/2
		citation := len(citations) - mid
		if citations[mid] >= citation {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return len(citations) - left
}
