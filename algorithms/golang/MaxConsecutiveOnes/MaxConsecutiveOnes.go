package MaxConsecutiveOnes

func findMaxConsecutiveOnes(nums []int) int {
	var tmp, max int

	for _, v := range nums {
		if v == 1 {
			tmp++
			if tmp > max {
				max = tmp
			}
		} else {
			tmp = 0
		}

	}

	return max
}
