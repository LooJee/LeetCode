package main

import "fmt"

/*
Given an array A of non-negative integers, return an array consisting of all the even elements of A, followed by all the odd elements of A.
You may return any answer array that satisfies this condition.

Example 1:
Input: [3,1,2,4]
Output: [2,4,3,1]
The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.

Note:
1 <= A.length <= 5000
0 <= A[i] <= 5000
*/

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}

func sortArrayByParity(A []int) []int {
	i := 0
	j := len(A) - 1

	for i != j {
		for ; A[i]%2 == 0 && i != j; i++ {

		}

		for ; A[j]%2 == 1 && i != j; j-- {

		}

		A[i], A[j] = A[j], A[i]
	}

	return A
}
