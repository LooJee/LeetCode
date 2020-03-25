package main

/*
Given a fixed length array arr of integers, duplicate each occurrence of zero, shifting the remaining elements to the right.
Note that elements beyond the length of the original array are not written.
Do the above modifications to the input array in place, do not return anything from your function.

Example 1:
Input: [1,0,2,3,0,4,5,0]
Output: null
Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]

Example 2:
Input: [1,2,3]
Output: null
Explanation: After calling your function, the input array is modified to: [1,2,3]

Note:
1 <= arr.length <= 10000
0 <= arr[i] <= 9
*/

func main() {
	duplicateZeros([]int{1, 0, 2, 3, 0, 4, 5, 0})
	duplicateZeros([]int{1, 0, 0, 0, 2, 3})
	duplicateZeros([]int{0, 1, 7, 6, 0, 2, 0, 7})
}

func duplicateZeros(arr []int) {
	zCnt := 0

	for i := 0; i < len(arr)-zCnt; i++ {
		if arr[i] == 0 {
			zCnt++
		}
	}

	for i := len(arr) - 1; i > 0; i-- {
		if arr[i-zCnt] != 0 {
			arr[i] = arr[i-zCnt]
		} else {
			arr[i] = 0
			i--
			arr[i] = 0
		}
	}
}
