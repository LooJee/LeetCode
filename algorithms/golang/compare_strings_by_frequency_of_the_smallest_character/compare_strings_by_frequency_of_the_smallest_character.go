package main

import "fmt"

/*
Let's define a function f(s) over a non-empty string s, which calculates the frequency of the smallest character in s.
For example, if s = "dcce" then f(s) = 2 because the smallest character is "c" and its frequency is 2.
Now, given string arrays queries and words, return an integer array answer, where each answer[i] is the number of words such that f(queries[i]) < f(W), where W is a word in words.

Example 1:
Input: queries = ["cbd"], words = ["zaaaz"]
Output: [1]
Explanation: On the first query we have f("cbd") = 1, f("zaaaz") = 3 so f("cbd") < f("zaaaz").

Example 2:
Input: queries = ["bbb","cc"], words = ["a","aa","aaa","aaaa"]
Output: [1,2]
Explanation: On the first query only f("bbb") < f("aaaa"). On the second query both f("aaa") and f("aaaa") are both > f("cc").

Constraints:
1 <= queries.length <= 2000
1 <= words.length <= 2000
1 <= queries[i].length, words[i].length <= 10
queries[i][j], words[i][j] are English lowercase letters.
["bba","abaaaaaa","aaaaaa","bbabbabaab","aba","aa","baab","bbbbbb","aab","bbabbaabb"]
["aaabbb","aab","babbab","babbbb","b","bbbbbbbbab","a","bbbbbbbbbb","baaabbaab","aa"]
*/

func main() {
	fmt.Println(numSmallerByFrequency([]string{"cbd"}, []string{"zaaaz"}))
	fmt.Println(numSmallerByFrequency([]string{"bbb", "cc"}, []string{"a", "aa", "aaa", "aaaa"}))
	fmt.Println(numSmallerByFrequency([]string{"bba", "abaaaaaa", "aaaaaa", "bbabbabaab", "aba", "aa", "baab", "bbbbbb", "aab", "bbabbaabb"},
		[]string{"aaabbb", "aab", "babbab", "babbbb", "b", "bbbbbbbbab", "a", "bbbbbbbbbb", "baaabbaab", "aa"}))
}

func numSmallerByFrequency(queries []string, words []string) []int {
	ret := make([]int, len(queries))
	smallest := make([]int, 11)

	for _, word := range words {
		smallestCh := 'z' + 1
		smallestSize := 0

		for _, ch := range word {
			if ch < smallestCh {
				smallestCh = ch
				smallestSize = 1
			} else if ch == smallestCh {
				smallestSize++
			}
		}
		smallest[smallestSize]++
	}

	for idx, query := range queries {
		smallestCh := 'z' + 1
		smallestSize := 0

		for _, ch := range query {
			if ch < smallestCh {
				smallestCh = ch
				smallestSize = 1
			} else if ch == smallestCh {
				smallestSize++
			}
		}

		for i := smallestSize + 1; i <= 10; i++ {
			if smallest[i] != 0 {
				ret[idx] += smallest[i]
			}
		}
	}

	return ret
}
