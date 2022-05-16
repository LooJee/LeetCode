package lcof_fib

func fib(n int) int {
	if n < 2 {
		return n
	}
	mod := int(1e9 + 7)
	a := 0
	b := 1
	for i := 2; i <= n; i++ {
		a, b = b, (a+b)%mod
	}

	return b
}
