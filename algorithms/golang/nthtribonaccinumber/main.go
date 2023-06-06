package nthtribonaccinumber

func tribonacci(n int) int {
	var (
		t = []int{0, 1, 1}
	)

	if n <= 2 {
		return t[n]
	}

	for i := 3; i <= n; i++ {
		t[0], t[1], t[2] = t[1], t[2], t[0]+t[1]+t[2]
	}

	return t[2]
}
