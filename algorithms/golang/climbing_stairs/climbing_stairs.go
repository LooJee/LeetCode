package climbing_stairs

func climbStairs(n int) int {
	var (
		a, b = 1, 2
	)

	if n <= 2 {
		return n
	}

	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}
