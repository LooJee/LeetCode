package lcof_min_array

func minArray(numbers []int) int {
	l, r := 0, len(numbers)-1
	var mid int
	for l <= r {
		mid = (l + r) / 2
		if numbers[mid] < numbers[r] {
			//右边都比mid大，最小值在左边，可以忽略
			r = mid
		} else if numbers[mid] > numbers[r] {
			//右边的比mid小，可以忽略左边的
			l = mid + 1
		} else {
			r--
		}
	}

	return numbers[mid]
}
