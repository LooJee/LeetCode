package four_sum

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}

	sort.Ints(nums)

	results := make([][]int, 0)

	a := 0

	for a < len(nums)-3 {
		b := a + 1
		for b < len(nums)-2 {
			c, d := b+1, len(nums)-1

			for c < d {
				sum := nums[a] + nums[b] + nums[c] + nums[d]

				if sum <= target {
					if sum == target {
						results = append(results, []int{nums[a], nums[b], nums[c], nums[d]})
					}

					c++
					for nums[c] == nums[c-1] && c < d {
						c++
					}
				} else if sum > target {
					d--
					for nums[d] == nums[d+1] && c < d {
						d--
					}
				}
			}

			b++
			for nums[b] == nums[b-1] && b < len(nums)-2 {
				b++
			}
		}

		a++
		for nums[a] == nums[a-1] && a < len(nums)-3 {
			a++
		}
	}

	return results
}
