package sort_colors

/*
使用快排
*/
func sortColors(nums []int) {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	for k := i; k < len(nums); k++ {
		if nums[k] == 1 {
			nums[i], nums[k] = nums[k], nums[i]
			i++
		}
	}
}
