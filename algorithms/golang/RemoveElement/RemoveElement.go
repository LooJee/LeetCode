package RemoveElement

func removeElement(nums []int, val int) int {
	var (
		j = len(nums) - 1
		i = 0
	)

	for i <= j {
		if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		} else {
			i++
		}
	}

	return j + 1
}
