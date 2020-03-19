package main

import "fmt"

/*
Given an unsorted array of integers, find the length of longest continuous increasing subsequence (subarray).

Example 1:
Input: [1,3,5,4,7]
Output: 3
Explanation: The longest continuous increasing subsequence is [1,3,5], its length is 3.
Even though [1,3,5,7] is also an increasing subsequence, it's not a continuous one where 5 and 7 are separated by 4.
Example 2:
Input: [2,2,2,2,2]
Output: 1
Explanation: The longest continuous increasing subsequence is [2], its length is 1.
Note: Length of the array will not exceed 10,000.
*/

func main() {
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))
	fmt.Println(findLengthOfLCIS([]int{2, 2, 2, 2, 2}))
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 7}))
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 2, 3, 4, 5}))
	fmt.Println(findLengthOfLCIS([]int{}))
}

func findLengthOfLCIS(nums []int) int {
	lcis := 0
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if cnt == 0 || nums[i-1] < nums[i] {
			cnt++
		} else {
			cnt = 1
		}

		if cnt > lcis {
			lcis = cnt
		}
	}

	return lcis
}
