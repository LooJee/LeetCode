package lcof_exchange

func exchange(nums []int) []int {
	l, r := 0, len(nums)-1

	for l < r {
		if nums[l]%2 == 1 {
			l++
			continue
		} else {
			for l <= r && nums[r]%2 == 0 {
				r--
			}
			if r <= l {
				break
			}
			nums[l], nums[r] = nums[r], nums[l]
		}
	}

	return nums
}
