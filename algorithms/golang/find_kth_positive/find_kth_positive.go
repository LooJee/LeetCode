package find_kth_positive

//因为输入数组是一个有序的数组，所以可以采用二分法，
//如果arr[i] == i，那么i之前的数据是没有缺失的，数据区间往右移
//如果arr[i] > i,那么[0:1]之间缺失的元素个数为n = arr[i]-i个
//如果n > k，那么需要将区间往左移，否则，将区间往右移
func findKthPositive(arr []int, k int) int {
	if arr[0] > k {
		return k
	}
	l, r := 0, len(arr)-1
	var mid int
	for l <= r {
		mid = (l + r) >> 1
		if arr[mid]-mid-1 >= k {
			if arr[mid-1]-(mid-1)-1 < k {
				mid = mid - 1
				break
			}
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return arr[mid] + (k - (arr[mid] - mid - 1))
}
