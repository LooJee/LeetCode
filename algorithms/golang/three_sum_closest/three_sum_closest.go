package three_sum_closest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}

	sort.Ints(nums)

	pDiff := math.MaxInt
	sum := 0
	i := 0
	j := i + 1
	k := len(nums) - 1

	for i < len(nums)-2 {
		for j < k {
			tmp := nums[i] + nums[j] + nums[k]
			if tmp == target {
				return tmp
			}

			diff := int(math.Abs(float64(target - tmp)))
			if diff < pDiff {
				sum = tmp
				pDiff = diff
			}

			if tmp < target {
				j++
				for nums[j] == nums[j-1] && j < k {
					j++
				}
			} else {
				k--
				for nums[k] == nums[k+1] && j < k {
					k--
				}
			}
		}
		i++
		j = i + 1
		k = len(nums) - 1
	}

	return sum
}
