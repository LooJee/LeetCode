package FibonacciNumber

func fib(n int) int {
	if n == 0 {
		return 0
	}

	// 容量比 n 大 1， 避免 n 为 1 时初始化 dp[1] 发生数组越界
	dp := make([]int, n+1)

	dp[0], dp[1] = 0, 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
