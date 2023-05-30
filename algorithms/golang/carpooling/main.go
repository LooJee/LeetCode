package carpooling

import "fmt"

func carPooling(trips [][]int, capacity int) bool {
	var (
		nums = make([]int, 1000)
		diff = NewDifference(nums)
	)

	for _, trip := range trips {
		diff.Increment(trip[1], trip[2]-1, trip[0])
	}

	nums = diff.Result()
	fmt.Println(nums)

	for _, num := range nums {
		if num > capacity {
			return false
		}
	}

	return true
}

type Difference struct {
	diff []int
}

func NewDifference(nums []int) Difference {
	diff := Difference{diff: make([]int, len(nums))}

	for i := 1; i < len(nums); i++ {
		diff.diff[i] = nums[i] - nums[i-1]
	}

	return diff
}

func (diff Difference) Increment(i, j, val int) {
	diff.diff[i] += val
	if j+1 < len(diff.diff) {
		diff.diff[j+1] -= val
	}

	fmt.Println(diff.diff)
}

func (diff Difference) Result() []int {
	nums := make([]int, len(diff.diff))

	nums[0] = diff.diff[0]
	for i := 1; i < len(diff.diff); i++ {
		nums[i] = nums[i-1] + diff.diff[i]
	}

	return nums
}
