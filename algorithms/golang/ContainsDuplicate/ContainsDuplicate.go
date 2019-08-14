package ContainsDuplicate

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})

	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = struct{}{}
	}

	return false
}
