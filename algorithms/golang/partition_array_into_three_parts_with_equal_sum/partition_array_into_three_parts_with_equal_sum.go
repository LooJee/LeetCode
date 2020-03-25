package main

import "fmt"

/*
Given an array A of integers, return true if and only if we can partition the array into three non-empty parts with equal sums.
Formally, we can partition the array if we can find indexes i+1 < j with (A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ... + A[j-1] == A[j] + A[j-1] + ... + A[A.length - 1])

Example 1:
Input: A = [0,2,1,-6,6,-7,9,1,2,0,1]
Output: true
Explanation: 0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1

Example 2:
Input: A = [0,2,1,-6,6,7,9,-1,2,0,1]
Output: false

Example 3:
Input: A = [3,3,6,5,-2,2,5,1,-9,4]
Output: true
Explanation: 3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4

Constraints:
3 <= A.length <= 50000
-10^4 <= A[i] <= 10^4
*/

func main() {
	fmt.Println(canThreePartsEqualSum([]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}))
	fmt.Println(canThreePartsEqualSum([]int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1}))
	fmt.Println(canThreePartsEqualSum([]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}))
}

func canThreePartsEqualSum(A []int) bool {
	var sum int

	for _, v := range A {
		sum += v
	}

	if sum % 3 != 0 {
		return false
	}

	oneThird := sum / 3
	var tmpSum int
	var sumCnt int
	for idx, v := range A {
		tmpSum += v
		if tmpSum == oneThird {
			sumCnt++
			tmpSum = 0
		}

		if sumCnt == 2 && idx < len(A)-1 {
			return true
		}
	}

	return false
}
