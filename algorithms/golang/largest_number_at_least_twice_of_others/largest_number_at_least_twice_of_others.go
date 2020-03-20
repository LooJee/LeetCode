package main

import (
	"fmt"
	"math"
)

/*
In a given integer array nums, there is always exactly one largest element.

Find whether the largest element in the array is at least twice as much as every other number in the array.

If it is, return the index of the largest element, otherwise return -1.

Example 1:

Input: nums = [3, 6, 1, 0]
Output: 1
Explanation: 6 is the largest integer, and for every other number in the array x,
6 is more than twice as big as x.  The index of value 6 is 1, so we return 1.


Example 2:

Input: nums = [1, 2, 3, 4]
Output: -1
Explanation: 4 isn't at least as big as twice the value of 3, so we return -1.


Note:

nums will have a length in the range [1, 50].
Every nums[i] will be an integer in the range [0, 99].
*/

func main() {
	fmt.Println(dominantIndex([]int{3, 6, 1, 0}))
	fmt.Println(dominantIndex([]int{1, 2, 3, 4}))
	fmt.Println(dominantIndex([]int{0, 2, 3, 0}))
	fmt.Println(dominantIndex([]int{0}))
}

func dominantIndex(nums []int) int {
	max1 := math.MinInt32
	idx1 := -1
	max2 := math.MinInt32

	for i := 0; i < len(nums); i++ {
		if nums[i] > max1 {
			max1, max2 = nums[i], max1
			idx1 = i
		} else if nums[i] > max2 {
			max2 = nums[i]
		}
	}

	if max1 >= max2*2 {
		return idx1
	}

	return -1
}
