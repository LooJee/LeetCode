package main

import "fmt"

/*
Given a list of dominoes, dominoes[i] = [a, b] is equivalent to dominoes[j] = [c, d] if and only if either (a==c and b==d), or (a==d and b==c) - that is, one domino can be rotated to be equal to another domino.
Return the number of pairs (i, j) for which 0 <= i < j < dominoes.length, and dominoes[i] is equivalent to dominoes[j].

Example 1:
Input: dominoes = [[1,2],[2,1],[3,4],[5,6]]
Output: 1

Constraints:
1 <= dominoes.length <= 40000
1 <= dominoes[i][j] <= 9
*/

func main() {
	fmt.Println(numEquivDominoPairs([][]int{
		{1, 2}, {2, 1}, {3, 4}, {5, 6},
	}))
	fmt.Println(numEquivDominoPairs([][]int{
		{1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2},
	}))
}

func numEquivDominoPairs(dominoes [][]int) int {
	dM := make(map[int]int, len(dominoes))
	ret := 0

	for _, subD := range dominoes {
		if subD[0] < subD[1] {
			dM[subD[0]*10+subD[1]]++
		} else {
			dM[subD[1]*10+subD[0]]++
		}
	}

	for _, val := range dM {
		if val > 1 {
			ret += val * (val - 1) / 2
		}
	}

	return ret
}
