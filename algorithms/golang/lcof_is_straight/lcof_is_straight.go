package lcof_is_straight

func isStraight(nums []int) bool {
	quickSort(nums, 0, len(nums)-1)

	joker := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			joker++
			continue
		}

		if i+1 < len(nums) {
			if nums[i+1]-nums[i] == 1 {
				continue
			}
			if nums[i+1]-nums[i] == 0 || nums[i+1]-nums[i] > joker+1 {
				return false
			} else {
				joker -= nums[i+1] - nums[i] + 1
			}
		}
	}

	return true
}

func quickSort(nums []int, s, e int) {
	if s >= e {
		return
	}

	p := partition(nums, s, e)
	quickSort(nums, s, p-1)
	quickSort(nums, p+1, e)
}

func partition(nums []int, s, p int) int {
	i := s
	j := s

	for i < p {
		if nums[i] < nums[p] {
			if j != i {
				nums[j], nums[i] = nums[i], nums[j]
			}
			j++
		}
		i++
	}

	nums[j], nums[p] = nums[p], nums[j]

	return j
}
