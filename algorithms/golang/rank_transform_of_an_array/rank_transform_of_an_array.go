package main

import (
	"fmt"
	"sort"
)

/*
Given an array of integers arr, replace each element with its rank.
The rank represents how large the element is. The rank has the following rules:
Rank is an integer starting from 1.
The larger the element, the larger the rank. If two elements are equal, their rank must be the same.
Rank should be as small as possible.

Example 1:
Input: arr = [40,10,20,30]
Output: [4,1,2,3]
Explanation: 40 is the largest element. 10 is the smallest. 20 is the second smallest. 30 is the third smallest.

Example 2:
Input: arr = [100,100,100]
Output: [1,1,1]
Explanation: Same elements share the same rank.

Example 3:
Input: arr = [37,12,28,9,100,56,80,5,12]
Output: [5,3,4,2,8,6,7,1,3]

Constraints:
0 <= arr.length <= 10^5
-10^9 <= arr[i] <= 10^9
*/

func main() {
	fmt.Println(arrayRankTransform([]int{40, 10, 20, 30}))
	fmt.Println(arrayRankTransform([]int{100, 100, 100}))
	fmt.Println(arrayRankTransform([]int{37, 12, 28, 9, 100, 56, 80, 5, 12}))
}

func arrayRankTransform(arr []int) []int {
	ret := make([]int, 0, len(arr))
	dict := make(map[int]int, len(arr))

	ret = append(ret, arr...)
	sort.Ints(ret)
	rank := 1
	for idx, val := range ret {
		if v, ok := dict[val]; ok && v > idx+1 || !ok {
			dict[val] = rank
			rank++
		}
	}

	for idx, val := range arr {
		ret[idx] = dict[val]
	}

	return ret
}
