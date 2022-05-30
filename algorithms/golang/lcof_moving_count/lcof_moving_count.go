package lcof_moving_count

func movingCount(m int, n int, k int) int {
	moved := make(map[int]map[int]struct{})

	return count(moved, 0, 0, m, n, k)
}

func count(moved map[int]map[int]struct{}, i, j, m, n, k int) int {
	if i > m-1 || j > n-1 || digitSum(i)+digitSum(j) > k || hasMoved(moved, i, j) {
		return 0
	}

	move(moved, i, j)

	return 1 + count(moved, i+1, j, m, n, k) + count(moved, i, j+1, m, n, k)
}

func digitSum(num int) int {
	return num/10 + num%10
}

func hasMoved(moved map[int]map[int]struct{}, i, j int) bool {
	if m, ok := moved[i]; ok {
		if _, ok := m[j]; ok {
			return true
		}
	}

	return false
}

func move(moved map[int]map[int]struct{}, i, j int) map[int]map[int]struct{} {
	m, ok := moved[i]
	if !ok {
		m = map[int]struct{}{}
	}

	m[j] = struct{}{}
	moved[i] = m
	return moved
}
