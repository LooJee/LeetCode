package FindAllNumbersDisappearedinanArray

func findDisappearedNumbers(nums []int) []int {
	other := make([]int, len(nums)+1)

	for _, v := range nums {
		other[v]++
	}

	ret := make([]int, 0, len(nums))
	for k, v := range other {
		if v == 0 && k != 0 {
			ret = append(ret, k)
		}
	}

	return ret
}
