package summary_ranges

import "testing"

func compareStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if b[i] != v {
			return false
		}
	}

	return true
}

func TestSummaryRanges(t *testing.T) {
	if !compareStringSlice(summaryRanges([]int{0, 1, 2, 4, 5, 7}), []string{"0->2", "4->5", "7"}) {
		t.Fatalf("should equal, : %v", summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	}
}
