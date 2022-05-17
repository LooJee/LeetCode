package lcof_max_value

import "testing"

func TestMaxValue(t *testing.T) {
	t.Log(maxValue([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))

	t.Log(maxValue([][]int{
		{1, 2, 5},
		{3, 2, 1},
	}))
}
