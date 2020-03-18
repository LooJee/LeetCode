package main

import "fmt"

func main() {
	fmt.Println(canPlaceFlowers([]int{0, 0, 1, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{0, 1, 0, 1, 0, 0}, 1))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	tCount := 0

	for i, v := range flowerbed {
		if v == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			tCount++
		}
	}

	return n <= tCount
}
