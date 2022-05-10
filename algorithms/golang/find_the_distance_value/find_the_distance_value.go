package find_the_distance_value

import "math"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sum := 0
	for _, i := range arr1 {
		ok := true
		for _, j := range arr2 {
			if int(math.Abs(float64(i-j))) <= d {
				ok = false
				break
			}
		}

		if ok {
			sum++
		}
	}

	return sum
}
