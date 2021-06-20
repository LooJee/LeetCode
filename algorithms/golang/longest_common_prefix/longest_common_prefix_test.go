package longest_common_prefix

import "testing"

func TestLongCommonPrefix(t *testing.T) {
	if longestCommonPrefix([]string{"flower", "flow", "flight"}) != "fl" {
		t.Fatal("longest common prefix should be \"fl\"")
	}

	if longestCommonPrefix([]string{}) != "" {
		t.Fatal("longest common prefix should be \"\"")
	}
}
