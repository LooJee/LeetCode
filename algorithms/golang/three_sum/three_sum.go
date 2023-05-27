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
		remain := 0 - nums[i]
		for j < k {
			if nums[j] == nums[j-1] && j-1 != i {
				j++
				continue
			}

			if nums[j]+nums[k] == remain {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
			} else if nums[j]+nums[k] < remain {
				j++
			} else if nums[j]+nums[k] > remain {
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
