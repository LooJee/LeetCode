package valid_parentheses

import "testing"

func TestIsValid(t *testing.T) {
	if isValid("{") {
		t.Fatal("\"{\" is not valid")
	}

	if !isValid("{}") {
		t.Fatal("\"{}\" is valid")
	}

	if !isValid("()[]{}") {
		t.Fatal("\"()[]{}\" is valid")
	}

	if isValid("([)]") {
		t.Fatal("\"([)]\" is not valid")
	}

}
