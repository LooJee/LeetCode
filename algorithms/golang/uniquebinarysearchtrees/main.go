package uniquebinarysearchtrees

func numTrees(n int) int {
	memos := make([][]int, n+1)

	for i := 1; i <= n; i++ {
		memos[i] = make([]int, n+1)
	}

	return count(memos, 1, n)
}

func count(memos [][]int, left, right int) int {
	if left > right {
		return 1
	}

	if cnt := memos[left][right]; cnt != 0 {
		return cnt
	}

	var cnt int

	for i := left; i <= right; i++ {
		leftCnt := count(memos, left, i-1)
		rightCnt := count(memos, i+1, right)
		cnt += leftCnt * rightCnt
	}

	memos[left][right] = cnt

	return cnt
}
