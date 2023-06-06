package mincostclimbingstairs

import "math"

func minCostClimbingStairs(cost []int) int {
	var (
		a, b = cost[0], 0
	)

	for i := 1; i < len(cost); i++ {
		a, b = int(math.Min(float64(a), float64(b)))+cost[i], a
	}

	return int(math.Min(float64(a), float64(b)))
}
