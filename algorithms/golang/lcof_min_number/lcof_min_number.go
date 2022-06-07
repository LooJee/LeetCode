package lcof_min_number

import (
	"strconv"
	"strings"
)

func minNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}

	numsStrs := make([]string, len(nums))

	for i := 0; i < len(nums); i++ {
		numsStrs[i] = strconv.Itoa(nums[i])
	}

	quickSort(numsStrs, 0, len(numsStrs)-1)

	return strings.Join(numsStrs, "")
}

func quickSort(numStrs []string, s, e int) {
	if s >= e {
		return
	}

	p := partition(numStrs, s, e)

	quickSort(numStrs, s, p-1)
	quickSort(numStrs, p+1, e)
}

func partition(numStrs []string, s, p int) int {
	i := s
	j := s

	for i < p {
		if numStrs[i]+numStrs[p] < numStrs[p]+numStrs[i] {
			if j != i {
				numStrs[i], numStrs[j] = numStrs[j], numStrs[i]
			}
			j++
		}
		i++
	}

	numStrs[j], numStrs[p] = numStrs[p], numStrs[j]

	return j
}
