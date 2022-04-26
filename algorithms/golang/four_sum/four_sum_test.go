package four_sum

import (
	"testing"
)

func TestFourSum(t *testing.T) {
	t.Log(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	t.Log(fourSum([]int{2, 2, 2, 2, 2}, 8))
}
