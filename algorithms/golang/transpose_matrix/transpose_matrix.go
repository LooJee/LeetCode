package main

import "fmt"

/*
Given a matrix A, return the transpose of A.
The transpose of a matrix is the matrix flipped over it's main diagonal, switching the row and column indices of the matrix.

Example 1:

Input: [[1,2,3],[4,5,6],[7,8,9]]
Output: [[1,4,7],[2,5,8],[3,6,9]]
Example 2:

Input: [[1,2,3],[4,5,6]]
Output: [[1,4],[2,5],[3,6]]
*/

func main() {
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}

func transpose(A [][]int) [][]int {
	B := make([][]int, 0, len(A[0]))

	for i := 0; i < len(A[0]); i++ {
		a := make([]int, len(A))
		for j := 0; j < len(A); j++ {
			a[j] = A[j][i]
		}
		B = append(B, a)
	}

	return B
}
