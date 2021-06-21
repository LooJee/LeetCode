package sqrtx

import "math"

func mySqrt(x int) int {
	if x <= 0 {
		return 0
	}
	num := float64(x)

	for {
		prev := num
		num = (prev + float64(x)/prev) / 2
		if math.Abs(num-prev) < 0.1 {
			break
		}
	}

	return int(num)
}
