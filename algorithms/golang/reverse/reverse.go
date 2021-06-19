package reverse

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	xRune := []rune(strconv.Itoa(x))

	positive := true

	if xRune[0] == '-' {
		positive = false
		xRune = xRune[1:]
	}

	runeLength := len(xRune)

	for i := 0; i < runeLength/2; i++ {
		xRune[i], xRune[runeLength-i-1] = xRune[runeLength-i-1], xRune[i]
	}

	reverseI, _ := strconv.Atoi(string(xRune))

	if !positive {
		reverseI = reverseI * -1
	}

	if reverseI < math.MinInt32 || reverseI > math.MaxInt32 {
		return 0
	}

	return reverseI
}

func reverse2(x int) int {
	num := 0

	for x != 0 {
		num = num*10 + x%10
		x /= 10
	}

	if num < math.MinInt32 || num > math.MaxInt32 {
		return 0
	}

	return num
}
