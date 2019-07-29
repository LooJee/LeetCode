package TwoSum

func twoSum(nums []int, target int) []int {
	numLen := len(nums)
	numDict := make(map[int]int)

	for i := 0; i < numLen; i++ {
		if idx, ok := numDict[target-nums[i]]; ok {
			return []int{idx, i}
		} else {
			numDict[nums[i]] = i
		}
	}
	return nil
}
