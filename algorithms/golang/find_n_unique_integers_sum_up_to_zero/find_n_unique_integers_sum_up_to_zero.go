package main

import "fmt"

/*
Given an integer n, return any array containing n unique integers such that they add up to 0.

Example 1:
Input: n = 5
Output: [-7,-1,1,3,4]
Explanation: These arrays also are accepted [-5,-1,1,2,3] , [-3,-1,2,-2,4].

Example 2:
Input: n = 3
Output: [-1,0,1]

Example 3:
Input: n = 1
Output: [0]

Constraints:
1 <= n <= 1000
*/

func main() {
	fmt.Println(sumZero(5))
	fmt.Println(sumZero(1))
}

func sumZero(n int) []int {
	nums := make([]int, n)
	var i, j int
	middle := n / 2

	if n%2 == 1 {
		nums[middle] = 0
		i = middle - 1
		j = middle + 1
	} else {
		i = middle - 1
		j = middle
	}

	for i >= 0 && j < n {
		nums[i] = -j
		nums[j] = j

		i--
		j++
	}

	return nums
}
