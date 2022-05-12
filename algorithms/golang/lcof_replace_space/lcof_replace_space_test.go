package lcof_replace_space

import "testing"

func TestReplaceSpace(t *testing.T) {
	if replaceSpace("We are happy.") != "We%20are%20happy." {
		t.Fatal("wrong result")
	}
}
