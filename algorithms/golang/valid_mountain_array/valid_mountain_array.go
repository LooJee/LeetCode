package main

import "fmt"

/*
Given an array A of integers, return true if and only if it is a valid mountain array.
Recall that A is a mountain array if and only if:
A.length >= 3
There exists some i with 0 < i < A.length - 1 such that:
A[0] < A[1] < ... A[i-1] < A[i]
A[i] > A[i+1] > ... > A[A.length - 1]

Example 1:
Input: [2,1]
Output: false

Example 2:
Input: [3,5,5]
Output: false

Example 3:
Input: [0,3,2,1]
Output: true

Note:
0 <= A.length <= 10000
0 <= A[i] <= 10000
*/

func main() {
	fmt.Println(validMountainArray([]int{2, 1}))
	fmt.Println(validMountainArray([]int{3, 5, 5}))
	fmt.Println(validMountainArray([]int{0, 3, 2, 1}))
	fmt.Println(validMountainArray([]int{0, 1, 2, 3, 4, 8, 9, 10, 11, 12, 11}))
}

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}

	i := 0
	j := len(A) - 1
	iC := true
	jC := true

	for i != j && (iC || jC) {
		if A[i] < A[i+1] {
			i++
		} else {
			iC = false
		}

		if A[j] < A[j-1] {
			j--
		} else {
			jC = false
		}
	}

	return (i == j) && (i > 0 && i < len(A)-1)
}
