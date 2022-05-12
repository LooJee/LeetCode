package lcof_find_repeat_number

func findRepeatNumber(nums []int) int {
	m := make(map[int]struct{})

	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return nums[i]
		} else {
			m[nums[i]] = struct{}{}
		}
	}

	return -1
}
