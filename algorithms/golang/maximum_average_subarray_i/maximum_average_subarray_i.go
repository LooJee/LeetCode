package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMaxAverage([]int{1,12,-5,-6,50,3}, 4))
	fmt.Println(findMaxAverage([]int{5}, 1))
}

func findMaxAverage(nums []int, k int) float64 {
	max := 0

	for i := 0; i < k; i++ {
		max += nums[i]
	}

	tmp := max
	for i := k; i < len(nums); i++ {
		tmp = tmp-nums[i-k]+nums[i]
		if tmp > max {
			max = tmp
		}
	}

	return float64(max) / float64(k)
}
