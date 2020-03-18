package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maximumProduct([]int{1, 2, 3}))
	fmt.Println(maximumProduct([]int{1, 2, 3, 4}))
	fmt.Println(maximumProduct([]int{-1, -2, -3}))
	fmt.Println(maximumProduct([]int{-4, -3, -2, -1, 60}))
	fmt.Println(maximumProduct([]int{722, 634, -504, -379, 163, -613, -842, -578, 750, 951, -158, 30, -238, -392, -487, -797, -157, -374, 999, -5, -521, -879, -858, 382, 626, 803, -347, 903, -205, 57, -342, 186, -736, 17, 83, 726, -960, 343, -984, 937, -758, -122, 577, -595, -544, -559, 903, -183, 192, 825, 368, -674, 57, -959, 884, 29, -681, -339, 582, 969, -95, -455, -275, 205, -548, 79, 258, 35, 233, 203, 20, -936, 878, -868, -458, -882, 867, -664, -892, -687, 322, 844, -745, 447, -909, -586, 69, -88, 88, 445, -553, -666, 130, -640, -918, -7, -420, -368, 250, -786}))
}

func maximumProduct(nums []int) int {
	min1 := math.MaxInt32
	min2 := math.MaxInt32
	max1 := math.MinInt32
	max2 := math.MinInt32
	max3 := math.MinInt32

	for _, v := range nums {
		if v >= max1 {
			max1, max2, max3 = v, max1, max2
		} else if v >= max2 {
			max2, max3 = v, max2
		} else if v >= max3 {
			max3 = v
		}

		if v <= min1 {
			min1, min2 = v, min1
		} else if v <= min2 {
			min2 = v
		}
	}

	//两个负数相乘为正，所以需要判断两个最小的数为负数时将其与最大的数相乘和最大的三个数相乘取最大值
	return biggerProduct(min1*min2, max2*max3) * max1
}

func biggerProduct(p1, p2 int) int {
	if p1 > p2 {
		return p1
	} else {
		return p2
	}
}
