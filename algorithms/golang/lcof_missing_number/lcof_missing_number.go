package lcof_missing_number

func missingNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			return nums[i] - 1
		}
	}

	return len(nums)
}
