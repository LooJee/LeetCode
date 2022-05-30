package lcof_exist

import "testing"

func TestExist(t *testing.T) {
	t.Log(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))
}
