package advantageshuffle

import "sort"

type numWithIdx struct {
	Num int
	Idx int
}

func advantageCount(nums1 []int, nums2 []int) []int {
	var (
		result       = make([]int, len(nums1))
		nums2WithIdx = make([]numWithIdx, len(nums2))
	)

	for index, num := range nums2 {
		nums2WithIdx[index] = numWithIdx{Num: num, Idx: index}
	}

	sort.Ints(nums1)
	sort.Slice(nums2WithIdx, func(i, j int) bool {
		return nums2WithIdx[i].Num < nums2WithIdx[j].Num
	})

	var (
		l, r = 0, len(nums1) - 1
	)

	for i := len(nums2WithIdx) - 1; i >= 0; i-- {
		if nums1[r] > nums2WithIdx[i].Num {
			result[nums2WithIdx[i].Idx] = nums1[r]
			r--
		} else {
			result[nums2WithIdx[i].Idx] = nums1[l]
			l++
		}
	}

	return result
}
