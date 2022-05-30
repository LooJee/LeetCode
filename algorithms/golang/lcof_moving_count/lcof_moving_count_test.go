package lcof_moving_count

import "testing"

func TestMovingCount(t *testing.T) {
	t.Log(movingCount(2, 3, 1))
	t.Log(movingCount(3, 1, 0))
}
