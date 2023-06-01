package randompickwithweight

import (
	"math/rand"
	"time"
)

type Solution struct {
	preSum []int
	n      int
}

func Constructor(w []int) Solution {
	solution := Solution{preSum: make([]int, len(w)+1)}

	for i := 1; i < len(solution.preSum); i++ {
		solution.preSum[i] = solution.preSum[i-1] + w[i-1]
	}
	solution.n = solution.preSum[len(solution.preSum)-1]

	rand.Seed(time.Now().UnixMilli())

	return solution
}

func (this *Solution) PickIndex() int {
	var (
		randNum = random(this.n)
		left    = 1
		right   = len(this.preSum) - 1
	)

	for left <= right {
		mid := left + (right-left)/2
		if this.preSum[mid] < randNum {
			left = mid + 1
		} else if this.preSum[mid] >= randNum {
			right = mid - 1
		}
	}

	return left - 1
}

func random(n int) int {
	return rand.Intn(n) + 1
}
