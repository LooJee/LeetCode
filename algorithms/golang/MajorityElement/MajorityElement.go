package MajorityElement

func majorityElement(nums []int) int {
	m := make(map[int]int)
	var majority int

	for _, ele := range nums {
		v, ok := m[ele]
		if !ok {
			m[ele] = 1
		} else {
			m[ele] = v + 1
		}

		if v+1 > len(nums)/2 {
			majority = ele
			break
		}
	}

	return majority
}
