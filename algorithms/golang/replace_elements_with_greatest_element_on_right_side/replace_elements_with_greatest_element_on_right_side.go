package main

import "fmt"

/*
Given an array arr, replace every element in that array with the greatest element among the elements to its right, and replace the last element with -1.
After doing so, return the array.

Example 1:
Input: arr = [17,18,5,4,6,1]
Output: [18,6,6,6,1,-1]

Constraints:
1 <= arr.length <= 10^4
1 <= arr[i] <= 10^5
*/

func main() {
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1}))
}

func replaceElements(arr []int) []int {
	greatestRight := arr[len(arr)-1]
	arr[len(arr)-1] = -1

	for i := len(arr) - 2; i >= 0; i-- {
		tmp := arr[i]
		arr[i] = greatestRight
		if tmp > greatestRight {
			greatestRight = tmp
		}
	}

	return arr
}
