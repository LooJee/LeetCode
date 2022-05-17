package lcof_nums_way

import "testing"

func TestNumsWay(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(numWays(i))
	}
}
