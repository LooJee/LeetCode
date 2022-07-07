package lcof_sum_nums

var v = 0

func sumNums(n int) int {
	v = 0
	return sum(n)
}

func sum(n int) int {
	_ = n > 0 && sum(n-1) > 0

	v = v + n
	return v
}
