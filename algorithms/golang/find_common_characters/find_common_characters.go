package main

import "fmt"

/*
Given an array A of strings made only from lowercase letters,
return a list of all characters that show up in all strings within the list (including duplicates).
For example, if a character occurs 3 times in all strings but not 4 times,
you need to include that character three times in the final answer.
You may return the answer in any order.

Example 1:
Input: ["bella","label","roller"]
Output: ["e","l","l"]

Example 2:
Input: ["cool","lock","cook"]
Output: ["c","o"]

Note:
1 <= A.length <= 100
1 <= A[i].length <= 100
A[i][j] is a lowercase letter
*/

func main() {
	fmt.Println(commonChars([]string{"bella", "label", "roller"}))
	fmt.Println(commonChars([]string{"cool", "lock", "cook"}))
}

func commonChars(A []string) []string {
	dict := make([][26]int, len(A))
	var ret []string

	for i := 0; i < len(A); i++ {
		for _, v := range A[i] {
			dict[i][v-'a']++
		}
	}

	for _, v := range A[0] {
		bingo := true
		for i := 0; i < len(dict); i++ {
			if dict[i][v-'a'] == 0 {
				bingo = false
				break
			} else {
				dict[i][v-'a']--
			}
		}
		if bingo {
			ret = append(ret, fmt.Sprintf("%c", v))
		}
	}

	return ret
}
