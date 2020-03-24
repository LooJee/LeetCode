package main

import (
	"fmt"
)

/*
Given an array of integers A sorted in non-decreasing order, return an array of the squares of each number, also in sorted non-decreasing order.

Example 1:
Input: [-4,-1,0,3,10]
Output: [0,1,9,16,100]

Example 2:
Input: [-7,-3,2,3,11]
Output: [4,9,9,49,121]

Note:
1 <= A.length <= 10000
-10000 <= A[i] <= 10000
A is sorted in non-decreasing order.
*/

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
	fmt.Println(sortedSquares([]int{-4, 0, 1, 1}))
}

func sortedSquares(A []int) []int {
	B := make([]int, len(A))
	i := 0
	j := len(A) - 1
	split := 0

	for split < len(A) {
		if A[split] >= 0 {
			break
		}
		split++
	}

	k := len(A) - 1

	for ; i != split && j >= split; k-- {
		if abs(A[i]) > A[j] {
			B[k] = A[i] * A[i]
			i++
		} else {
			B[k] = A[j] * A[j]
			j--
		}
	}

	for i != split {
		B[k] = A[i] * A[i]
		k--
		i++
	}

	for j >= split {
		B[k] = A[j] * A[j]
		k--
		j--
	}

	return B
}

func abs(v int) int {
	if v >= 0 {
		return v
	} else {
		return -v
	}
}
