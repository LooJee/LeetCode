package maximum_subarray

func maxSubArray(nums []int) int {
	sum := nums[0]
	tmp := nums[0]

	for i := 1; i < len(nums); i++ {
		tmp += nums[i]

		if tmp < nums[i] {
			tmp = nums[i]
		}

		if sum < tmp {
			sum = tmp
		}
	}

	return sum
}
