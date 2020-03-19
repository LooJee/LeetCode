package ContainsDuplicateII

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if idx, ok := m[nums[i]]; !ok {
			m[nums[i]] = i
		} else if i-idx <= k {
			return true
		} else {
			m[nums[i]] = i
		}
	}

	return false
}
