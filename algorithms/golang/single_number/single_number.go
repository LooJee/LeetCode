package single_number

func singleNumber(nums []int) int {
	xor := nums[0]

	for i := 1; i < len(nums); i++ {
		xor ^= nums[i]
	}

	return xor
}
