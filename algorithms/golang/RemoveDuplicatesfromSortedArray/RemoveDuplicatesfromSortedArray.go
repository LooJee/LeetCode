package RemoveDuplicatesfromSortedArray

func removeDuplicates(nums []int) int {
	j := 0
	numLen := len(nums)
	if numLen == 0 {
		return 0
	}

	for i := 0; i < numLen; i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1
}
