package lcof_get_least_numbers

func getLeastNumbers(arr []int, k int) []int {
	if k > len(arr) {
		return nil
	}

	quickSort(arr, 0, len(arr)-1)

	return arr[:k]
}

func quickSort(nums []int, s, e int) {
	if s >= e {
		return
	}

	p := partition(nums, s, e)
	quickSort(nums, s, p-1)
	quickSort(nums, p+1, e)
}

func partition(nums []int, s, p int) int {
	i := s
	j := s

	for i < p {
		if nums[i] < nums[p] {
			if i != j {
				nums[j], nums[i] = nums[i], nums[j]
			}
			j++
		}
		i++
	}

	nums[p], nums[j] = nums[j], nums[p]

	return j
}
