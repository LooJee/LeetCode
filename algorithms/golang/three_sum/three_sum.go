package three_sum

import "sort"

func threeSum(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}

	sort.Ints(nums)

	i := 0
	j := 1
	k := len(nums) - 1

	for i < len(nums)-2 {
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum <= 0 {
				if sum == 0 {
					result = append(result, []int{nums[i], nums[j], nums[k]})
				}
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
		for nums[i] == nums[i-1] && i < len(nums)-2 {
			i++
		}
		j = i + 1
		k = len(nums) - 1
	}

	return result
}
