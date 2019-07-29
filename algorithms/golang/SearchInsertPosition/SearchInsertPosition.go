package SearchInsertPosition

func searchInsert(nums []int, target int) int {
	var i int
	for i = 0; i < len(nums); i++ {
		if nums[i] >= target {
			break
		}
	}
	return i
}
