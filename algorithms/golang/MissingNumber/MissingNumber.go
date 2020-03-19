package MissingNumber

func missingNumber(nums []int) int {
	n := len(nums)
	half := (n + 1) / 2
	sum := n * half
	if n%2 == 0 {
		sum += half
	}

	for i := 0; i < n; i++ {
		sum -= nums[i]
	}

	return sum
}
