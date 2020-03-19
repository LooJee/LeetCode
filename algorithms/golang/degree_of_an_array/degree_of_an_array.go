package main

import "fmt"

/*
Given a non-empty array of non-negative integers nums, the degree of this array is defined as the maximum frequency of any one of its elements.

Your task is to find the smallest possible length of a (contiguous) subarray of nums, that has the same degree as nums.

Example 1:
Input: [1, 2, 2, 3, 1]
Output: 2
Explanation:
The input array has a degree of 2 because both elements 1 and 2 appear twice.
Of the subarrays that have the same degree:
[1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
The shortest length is 2. So return 2.
Example 2:
Input: [1,2,2,3,1,4,2]
Output: 6

Note:
nums.length will be between 1 and 50,000.
nums[i] will be an integer between 0 and 49,999.
*/

func main() {
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1}))
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}))
	fmt.Println(findShortestSubArray([]int{2, 1, 1, 2, 1, 3, 3, 3, 1, 3, 1, 3, 2}))
}

func findShortestSubArray(nums []int) int {
	preMap := make(map[int]int) //previous position of the element
	lenMap := make(map[int]int) //longest length of the same elements
	dMap := make(map[int]int)   //degree map
	maxDegreeNum := -1
	maxDegree := 0

	for i := 0; i < len(nums); i++ {
		if v, ok := dMap[nums[i]]; !ok {
			preMap[nums[i]] = i
			lenMap[nums[i]] = 1
			dMap[nums[i]] = 1
		} else {
			pre, _ := preMap[nums[i]]
			preMap[nums[i]] = i
			lenI, _ := lenMap[nums[i]]
			lenMap[nums[i]] = lenI + i - pre
			dMap[nums[i]] = v + 1
		}

		if v, ok := dMap[nums[i]]; ok && v >= maxDegree {
			// compare the length between nums[i] and maxDegreeNum if they have same degree
			if v == maxDegree {
				lenMaxDegreeNum, _ := lenMap[maxDegreeNum]
				lenI, _ := lenMap[nums[i]]

				if lenI < lenMaxDegreeNum {
					maxDegreeNum = nums[i]
				}
			} else {
				maxDegree = v
				maxDegreeNum = nums[i]
			}
		}
	}

	//fmt.Println("max degree number", maxDegreeNum)
	shortest, _ := lenMap[maxDegreeNum]
	//fmt.Println(lenMap)
	//fmt.Println(preMap)
	//fmt.Println(dMap)

	return shortest
}
