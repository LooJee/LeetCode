package main

import "fmt"

/*
Given an array nums with n integers, your task is to check if it could become non-decreasing by modifying at most 1 element.

We define an array is non-decreasing if nums[i] <= nums[i + 1] holds for every i (0-based) such that (0 <= i <= n - 2).



Example 1:

Input: nums = [4,2,3]
Output: true
Explanation: You could modify the first 4 to 1 to get a non-decreasing array.
Example 2:

Input: nums = [4,2,1]
Output: false
Explanation: You can't get a non-decreasing array by modify at most one element.


Constraints:

1 <= n <= 10 ^ 4
- 10 ^ 5 <= nums[i] <= 10 ^ 5
*/

func main() {
	fmt.Println(checkPossibility([]int{4,2,3}))
	fmt.Println(checkPossibility([]int{4,2,1}))
	fmt.Println(checkPossibility([]int{3,4,2,3}))
	fmt.Println(checkPossibility([]int{-1,4,2,3}))
	fmt.Println(checkPossibility([]int{1,3,5,2,4}))
}

func checkPossibility(nums []int) bool {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if i != 0 && i != len(nums)-2 && nums[i] > nums[i+2] && nums[i-1] >= nums[i+1] {
				return false
			}
			count++
		}
	}

	return count <= 1
}
