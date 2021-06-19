package PalindromeNumber

import "testing"

func TestIsPalindrome(t *testing.T) {
	if !isPalindrome(121) {
		t.Fatal("121 is a palindrome")
	}

	if isPalindrome(-1) {
		t.Fatal("-1 not a palindrome")
	}

	if isPalindrome(123) {
		t.Fatal("123 not a palindrome")
	}

	if isPalindrome(10) {
		t.Fatal("10 not a palindrome")
	}

	if !isPalindrome(0) {
		t.Fatal("0 is a palindrome")
	}
}
