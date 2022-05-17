package arranging_coins

//这道题实际是一个等差数列求和的过程
// *
// * *
// * * *
// * * * *
//第k行有k个元素，如果k行都摆满，n=(k+1)*k/2
//我们需要求出k使得左值>=右值
func arrangeCoins(n int) int {
	l, r := 0, n
	k := 0
	for l < r {
		k = (l+r)/2 + 1
		if k*(k+1) <= n+n {
			l = k
		} else {
			r = k - 1
		}
	}

	return l
}
