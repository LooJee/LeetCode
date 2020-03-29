package main

import (
	"fmt"
	"math"
	"sort"
)

/*
Given an array of distinct integers arr, find all pairs of elements with the minimum absolute difference of any two elements.
Return a list of pairs in ascending order(with respect to pairs), each pair [a, b] follows
a, b are from arr
a < b
b - a equals to the minimum absolute difference of any two elements in arr

Example 1:
Input: arr = [4,2,1,3]
Output: [[1,2],[2,3],[3,4]]
Explanation: The minimum absolute difference is 1. List all pairs with difference equal to 1 in ascending order.

Example 2:
Input: arr = [1,3,6,10,15]
Output: [[1,3]]

Example 3:
Input: arr = [3,8,-10,23,19,-4,-14,27]
Output: [[-14,-10],[19,23],[23,27]]

Constraints:

2 <= arr.length <= 10^5
-10^6 <= arr[i] <= 10^6
*/

func main() {
	fmt.Println(minimumAbsDifference([]int{4, 2, 1, 3}))
	fmt.Println(minimumAbsDifference([]int{1, 3, 6, 10, 15}))
	fmt.Println(minimumAbsDifference([]int{3, 8, -10, 23, 19, -4, -14, 27}))
}

func minimumAbsDifference(arr []int) [][]int {
	minDist := math.MaxInt32
	ret := make([][]int, len(arr))
	sort.Ints(arr)

	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] < minDist {
			minDist = arr[i+1] - arr[i]
			ret = [][]int{{arr[i], arr[i+1]}}
		} else if arr[i+1]-arr[i] == minDist {
			ret = append(ret, []int{arr[i], arr[i+1]})
		}
	}

	return ret
}
