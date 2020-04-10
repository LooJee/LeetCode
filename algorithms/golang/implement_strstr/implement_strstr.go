package main

import (
	"fmt"
)

/*
Implement strStr().
Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:
Input: haystack = "hello", needle = "ll"
Output: 2

Example 2:
Input: haystack = "aaaaa", needle = "bba"
Output: -1
Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.
For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().
*/

func main() {
	fmt.Println(strStr("aaaaa", "bba"))
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("", ""))
	fmt.Println(strStr("a", "a"))
	fmt.Println(strStr("mississippi", "pi"))
}

/*BM 算法，使用坏字符原则*/
func strStr(haystack string, needle string) int {
	needleLen := len(needle)
	for i := 0; i <= len(haystack)-needleLen; {
		j := needleLen - 1
		for ; j >= 0; j-- {
			if needle[j] != haystack[i+j] {
				break
			}
		}

		if j < 0 {
			return i
		}

		k := j
		for ; k >= 0; k-- {
			if needle[k] == haystack[i+j] {
				break
			}
		}

		i += j - k
	}

	return -1
}
