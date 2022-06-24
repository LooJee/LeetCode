package lcof_median_finder

import "testing"

func TestFindMedian(t *testing.T) {
	mf := Constructor()

	mf.AddNum(-1)
	t.Log(mf.FindMedian())
	mf.AddNum(-2)
	t.Log(mf.FindMedian())
	mf.AddNum(-3)
	t.Log(mf.FindMedian())
	mf.AddNum(-4)
	t.Log(mf.FindMedian())
	mf.AddNum(-5)
	t.Log(mf.FindMedian())
}
