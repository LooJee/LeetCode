package main

import (
	"fmt"
	"math"
)

/*
Given a m * n matrix of distinct numbers, return all lucky numbers in the matrix in any order.
A lucky number is an element of the matrix such that it is the minimum element in its row and maximum in its column.

Example 1:
Input: matrix = [[3,7,8],[9,11,13],[15,16,17]]
Output: [15]
Explanation: 15 is the only lucky number since it is the minimum in its row and the maximum in its column

Example 2:
Input: matrix = [[1,10,4,2],[9,3,8,7],[15,16,17,12]]
Output: [12]
Explanation: 12 is the only lucky number since it is the minimum in its row and the maximum in its column.

Example 3:
Input: matrix = [[7,8],[1,2]]
Output: [7]

Constraints:
m == mat.length
n == mat[i].length
1 <= n, m <= 50
1 <= matrix[i][j] <= 10^5.
All elements in the matrix are distinct.
*/

func main() {
	fmt.Println(luckyNumbers([][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}))
	fmt.Println(luckyNumbers([][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}))
}

func luckyNumbers(matrix [][]int) []int {
	x := make([]int, len(matrix))    // minimum element in its row
	y := make([]int, len(matrix[0])) // maximum element in its column
	ret := []int{}

	for idx, _ := range x {
		x[idx] = math.MaxInt32
	}

	for idx, _ := range y {
		y[idx] = math.MinInt32
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] < x[i] {
				x[i] = matrix[i][j]
			}

			if matrix[i][j] > y[j] {
				y[j] = matrix[i][j]
			}
		}
	}

	for i := 0; i < len(x); i++ {
		for j := 0; j < len(y); j++ {
			if x[i] == y[j] {
				ret = append(ret, x[i])
			}
		}
	}

	return ret
}
