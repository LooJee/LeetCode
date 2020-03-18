package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findMaxAverage([]int{1,12,-5,-6,50,3}, 4))
	fmt.Println(findMaxAverage([]int{5}, 1))
}

func findMaxAverage(nums []int, k int) float64 {
	max := math.MinInt32

	for i := 0; i < len(nums) - k; i++ {
		sum := 0

		for j := i; j < k+i; j++ {
			sum += nums[j]
		}

		if sum > max {
			max = sum
		}
	}

	return float64(max) / float64(k)
}
