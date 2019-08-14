package KdiffPairsinanArray

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}

	m := make(map[int]int)
	pairs := 0

	for _, v := range nums {
		if cnt, ok := m[v]; ok {
			if k == 0 && cnt < 2 {
				pairs++
				m[v] = cnt+1
			}
			continue
		}

		if _, ok := m[v+k]; ok {
			pairs++
		}

		if _, ok := m[v-k]; ok {
			pairs++
		}

		m[v] = 1
	}

	return pairs
}
