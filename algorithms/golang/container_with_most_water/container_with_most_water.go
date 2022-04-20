package containerwithmostwater

import "math"

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	area := 0

	for i < j {
		min := 0
		k := j - i
		if height[i] < height[j] {
			min = height[i]
			i++
		} else {
			min = height[j]
			j--
		}

		area = int(math.Max(float64(area), float64(min*k)))
	}

	return area
}
