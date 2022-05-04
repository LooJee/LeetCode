package remove_duplicates_array

func removeDuplicates(nums []int) int {
	lenNum := len(nums)

	if lenNum < 2 {
		return lenNum
	}

	i := 2
	for j := 2; j < lenNum; j++ {
		if nums[j] != nums[i-2] {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}
