package lcof_nums_way

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n <= 2 {
		return n
	}

	mod := 1000000007
	a := 1
	b := 2
	for i := 3; i <= n; i++ {
		a, b = b, (a+b)%mod
	}

	return b
}
