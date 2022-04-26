package four_sum

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}

	sort.Ints(nums)

	results := make([][]int, 0)

	s := 0

	for s < len(nums)-3 {
		l := s + 1
		for l < len(nums)-2 {
			r := l + 1
			e := len(nums) - 1
			for r < e {
				sum := nums[s] + nums[l] + nums[r] + nums[e]
				if sum <= target {
					if sum == target {
						results = append(results, []int{nums[s], nums[l], nums[r], nums[e]})
					}
					r++
					for nums[r] == nums[r-1] && r < e {
						r++
					}
				} else {
					e--
					for nums[e] == nums[e+1] && r < e {
						e--
					}
				}
			}

			l++
			for nums[l] == nums[l-1] && l < len(nums)-2 {
				l++
			}
		}
		s++
		for nums[s] == nums[s-1] && s < len(nums)-3 {
			s++
		}
	}

	return results
}
