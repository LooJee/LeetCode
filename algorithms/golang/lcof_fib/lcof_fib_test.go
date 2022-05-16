package lcof_fib

import "testing"

func TestFib(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(fib(i))
	}
	t.Log(fib(1000000008))
	t.Log(fib(95))
}
