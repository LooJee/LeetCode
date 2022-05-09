package peak_index_in_mountain_array

func peakIndexInMountainArray(arr []int) int {
	l := 0
	r := len(arr) - 1
	var mid int

	for l <= r {
		mid = (l + r) / 2

		if mid-1 > 0 && arr[mid-1] > arr[mid] {
			r = mid - 1
		} else if mid+1 < len(arr) && arr[mid+1] > arr[mid] {
			l = mid + 1
		} else {
			return mid
		}
	}

	return 0
}
