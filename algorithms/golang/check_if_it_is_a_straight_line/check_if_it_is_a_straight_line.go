package main

import (
	"fmt"
	"math"
)

/*
You are given an array coordinates, coordinates[i] = [x, y], where [x, y] represents the coordinate of a point. Check if these points make a straight line in the XY plane.

Example 1:
Input: coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
Output: true

Example 2:
Input: coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
Output: false

Constraints:
2 <= coordinates.length <= 1000
coordinates[i].length == 2
-10^4 <= coordinates[i][0], coordinates[i][1] <= 10^4
coordinates contains no duplicate point.
*/

func main() {
	fmt.Println(checkStraightLine([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}))
	fmt.Println(checkStraightLine([][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}))
	fmt.Println(checkStraightLine([][]int{{-3, -2}, {-1, -2}, {2, -2}, {-2, -2}, {0, -2}}))
	fmt.Println(checkStraightLine([][]int{{-5, 1}, {-4, -1}, {-7, 4}, {-7, -7}, {5, -7}, {-3, 2}, {2, -5}}))
}

func checkStraightLine(coordinates [][]int) bool {
	for i := 1; i < len(coordinates)-1; i++ {
		sub1 := coordinates[i][0] - coordinates[i-1][0]
		if sub1 == 0 {
			sub1 = math.MaxInt32
		}

		sub2 := coordinates[i+1][0] - coordinates[i][0]
		if sub2 == 0 {
			sub2 = math.MaxInt32
		}
		if float64(coordinates[i][1]-coordinates[i-1][1])/float64(sub1) != float64(coordinates[i+1][1]-coordinates[i][1])/float64(sub2) {
			return false
		}
	}

	return true
}
