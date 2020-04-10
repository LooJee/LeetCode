package main

import "fmt"

/*
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.
Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:
Input: "A man, a plan, a canal: Panama"
Output: true

Example 2:
Input: "race a car"
Output: false
*/

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(""))
	fmt.Println(isPalindrome("0a"))
}

func isAlphaNumeric(c uint8) bool {
	if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
		return true
	}

	return false
}

func isPalindrome(s string) bool {
	lowerCase := 'A' - 'a'
	i := 0
	j := len(s) - 1
	for i <= j {
		vi := s[i]
		vj := s[j]

		if !isAlphaNumeric(vi) {
			i++
			continue
		}
		if !isAlphaNumeric(vj) {
			j--
			continue
		}

		if vi >= 'A' && vi <= 'Z' {
			vi -= uint8(lowerCase)
		}
		if vj >= 'A' && vj <= 'Z' {
			vj -= uint8(lowerCase)
		}

		if vi != vj {
			return false
		}

		i++
		j--
	}

	return true
}
