package main

import "fmt"

/*
Given an array nums of integers, return how many of them contain an even number of digits.

Example 1:
Input: nums = [12,345,2,6,7896]
Output: 2
Explanation:
12 contains 2 digits (even number of digits).
345 contains 3 digits (odd number of digits).
2 contains 1 digit (odd number of digits).
6 contains 1 digit (odd number of digits).
7896 contains 4 digits (even number of digits).
Therefore only 12 and 7896 contain an even number of digits.

Example 2:
Input: nums = [555,901,482,1771]
Output: 1
Explanation:
Only 1771 contains an even number of digits.

Constraints:
1 <= nums.length <= 500
1 <= nums[i] <= 10^5
*/

func main() {
	fmt.Println(findNumbers([]int{12, 345, 2, 6, 7896}))
	fmt.Println(findNumbers([]int{555, 901, 482, 1771}))
}

func findNumbers(nums []int) int {
	evenCnt := 0
	bases := []int{10, 100, 1000, 10000, 100000}

	for _, num := range nums {
		for idx, base := range bases {
			if num%base == num {
				if idx%2 == 1 {
					evenCnt++
				}
				break
			}
		}
	}

	return evenCnt
}
