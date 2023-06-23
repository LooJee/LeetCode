package maximumvalueofastringinanarray

import "strconv"

func maximumValue(strs []string) int {
	var (
		maxVal = 0
	)

	for _, str := range strs {
		num, err := strconv.Atoi(str)
		if err != nil {
			num = len(str)
		}

		if num > maxVal {
			maxVal = num
		}
	}

	return maxVal
}
