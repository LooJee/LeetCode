package lcof_exchange

import "testing"

func TestExchange(t *testing.T) {
	t.Log(exchange([]int{1, 2, 3, 4}))
	t.Log(exchange([]int{}))
	t.Log(exchange([]int{1}))
}
