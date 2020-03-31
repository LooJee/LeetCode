package main

import "fmt"

/*
Given a m * n matrix grid which is sorted in non-increasing order both row-wise and column-wise.
Return the number of negative numbers in grid.

Example 1:
Input: grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
Output: 8
Explanation: There are 8 negatives number in the matrix.

Example 2:
Input: grid = [[3,2],[1,0]]
Output: 0

Example 3:
Input: grid = [[1,-1],[-1,-1]]
Output: 3

Example 4:
Input: grid = [[-1]]
Output: 1

Constraints:
m == grid.length
n == grid[i].length
1 <= m, n <= 100
-100 <= grid[i][j] <= 100
*/

func main() {
	fmt.Println(countNegatives([][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}))
	fmt.Println(countNegatives([][]int{{3, 2}, {1, 0}}))
	fmt.Println(countNegatives([][]int{{1, -1}, {-1, -1}}))
	fmt.Println(countNegatives([][]int{{-1}}))
}

func countNegatives(grid [][]int) int {
	negatives := 0
	colLen := len(grid[0])

	for i := 0; i < len(grid); i++ {
		for j := 0; j < colLen; j++ {
			if grid[i][j] < 0 {
				negatives += (len(grid) - i) * (colLen - j)
				colLen = j
			}
		}
	}

	return negatives
}
