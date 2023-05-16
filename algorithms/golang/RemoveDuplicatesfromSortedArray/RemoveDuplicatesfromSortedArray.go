package RemoveDuplicatesfromSortedArray

func removeDuplicates(nums []int) int {
	var i = 0

	for j := 0; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}
