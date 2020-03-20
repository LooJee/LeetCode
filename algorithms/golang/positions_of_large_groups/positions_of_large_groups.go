package main

import "fmt"

/*
In a string S of lowercase letters, these letters form consecutive groups of the same character.
For example, a string like S = "abbxxxxzyy" has the groups "a", "bb", "xxxx", "z" and "yy".
Call a group large if it has 3 or more characters.  We would like the starting and ending positions of every large group.
The final answer should be in lexicographic order.

Example 1:

Input: "abbxxxxzzy"
Output: [[3,6]]
Explanation: "xxxx" is the single large group with starting  3 and ending positions 6.
Example 2:

Input: "abc"
Output: []
Explanation: We have "a","b" and "c" but no large group.
Example 3:

Input: "abcdddeeeeaabbbcd"
Output: [[3,5],[6,9],[12,14]]

Note:  1 <= S.length <= 1000
*/

func main() {
	fmt.Println(largeGroupPositions("abbxxxxzzy"))
	fmt.Println(largeGroupPositions("abcdddeeeeaabbbcd"))
	fmt.Println(largeGroupPositions("aaa"))
}

func largeGroupPositions(S string) [][]int {
	lg := make([][]int, 0)

	sIdx := 0
	for i := 1; i < len(S); i++ {
		if S[i] != S[sIdx] {
			if i-sIdx-1 >= 2 {
				lg = append(lg, []int{sIdx, i-1})
			}
			sIdx = i
		}
	}

	if S[len(S)-1] == S[sIdx] && len(S) - sIdx >= 3 {
		lg = append(lg, []int{sIdx, len(S)-1})
	}

	return lg
}
