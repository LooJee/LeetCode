package randompickwithblacklist

import (
	"math/rand"
	"sort"
	"time"
)

type Solution struct {
	idxMap        map[int]int
	whiteListSize int
}

func Constructor(n int, blacklist []int) Solution {
	var (
		solution = Solution{whiteListSize: n - len(blacklist), idxMap: map[int]int{}}
	)

	for index, element := range blacklist {
		solution.idxMap[element] = index
	}

	end := n - 1

	sort.Ints(blacklist)

	for i := 0; i < len(blacklist); i++ {
		// 判断 end 是否是 黑名单
		if _, ok := solution.idxMap[end]; ok {
			end--
			i--
			continue
		}
		solution.idxMap[blacklist[i]] = end

		end--
	}

	rand.Seed(time.Now().UnixMilli())

	return solution
}

func (this *Solution) Pick() int {
	random := rand.Intn(this.whiteListSize)
	if value, ok := this.idxMap[random]; ok {
		return value
	}

	return random
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
