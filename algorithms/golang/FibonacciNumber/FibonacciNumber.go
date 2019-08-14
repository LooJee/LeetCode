package FibonacciNumber

func fib(N int) int {
	if N == 0 || N == 1 {
		return N
	}

	a := 0
	b := 1
	for i := 1; i < N; i++ {
		a, b = b, a+b
	}

	return b
}
