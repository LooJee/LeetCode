package main

import "fmt"

/*
Given an integer array sorted in non-decreasing order, there is exactly one integer in the array that occurs more than 25% of the time.
Return that integer.

Example 1:
Input: arr = [1,2,2,6,6,6,6,7,10]
Output: 6

Constraints:

1 <= arr.length <= 10^4
0 <= arr[i] <= 10^5
*/

func main() {
	fmt.Println(findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 6, 7, 10}))
	fmt.Println(findSpecialInteger([]int{1, 1}))
	fmt.Println(findSpecialInteger([]int{1, 2, 3, 3}))
	fmt.Println(findSpecialInteger([]int{1}))
}

func findSpecialInteger(arr []int) int {
	leastTimes := len(arr) / 4

	dict := make([]int, arr[len(arr)-1]+1)

	for i := 0; i < len(arr); i++ {
		dict[arr[i]]++
		if dict[arr[i]] > leastTimes {
			return arr[i]
		}
	}

	return -1
}
