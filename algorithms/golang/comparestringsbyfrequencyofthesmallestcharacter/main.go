package comparestringsbyfrequencyofthesmallestcharacter

import (
	"math"
	"sort"
)

func numSmallerByFrequency(queries []string, words []string) []int {
	var (
		fwords = make([]int, len(words))
		result = make([]int, 0, len(queries))
	)

	for index, word := range words {
		fwords[index] = f(word)
	}

	sort.Ints(fwords)

	for _, query := range queries {
		fquery := f(query)

		var (
			l, r = 0, len(fwords) - 1
		)

		for l <= r {
			mid := l + (r-l)/2

			if fwords[mid] > fquery {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

		result = append(result, len(fwords)-l)
	}

	return result
}

func f(word string) int {
	var (
		minChar = math.MaxInt32
		cnt     = 0
	)

	for _, char := range word {
		if char < rune(minChar) {
			cnt = 1
			minChar = int(char)
		} else if char == rune(minChar) {
			cnt++
		}
	}

	return cnt
}
