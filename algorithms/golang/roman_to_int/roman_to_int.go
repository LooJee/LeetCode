package roman_to_int

func romanToInt(s string) int {
	romanMap := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	num := 0
	runeRoman := []rune(s)
	for i := 0; i < len(runeRoman); i++ {
		if i != len(runeRoman)-1 {
			if romanMap[runeRoman[i]] >= romanMap[runeRoman[i+1]] {
				num += romanMap[runeRoman[i]]
			} else {
				num += romanMap[runeRoman[i+1]] - romanMap[runeRoman[i]]
				i++
			}
		} else {
			num += romanMap[runeRoman[i]]
		}
	}

	return num
}
