package TwoSum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for idx, num := range nums {
		if v, ok := m[target-num]; ok && v != idx {
			return []int{idx, v}
		}

		m[num] = idx
	}

	return []int{}
}
